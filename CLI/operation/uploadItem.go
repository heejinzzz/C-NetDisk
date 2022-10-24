package operation

import (
	pb "C-NetDisk/CLI/proto"
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	pt "path"
)

func uploadFile(name string, uploadPath string) (chan int64, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	stream, err := CNetDiskClient.UploadFile(context.Background())
	if err != nil {
		return nil, err
	}
	ch := make(chan int64, DeliverSizeChannelLength)

	// send file data to server
	go func() {
		defer func() {
			err = file.Close()
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}()

		for {
			buf := make([]byte, ReadBytesEachTime)
			n, err := reader.Read(buf)
			if err == io.EOF {
				err = stream.CloseSend()
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				return
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			err = stream.Send(&pb.UploadFileRequest{Name: uploadPath + "/" + pt.Base(name), Data: buf[:n]})
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
	}()

	// get server's response
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				ch <- -1
				return
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			if res.ResSignal == pb.Signal_FinishReceive {
				ch <- -1
				return
			}
			ch <- res.GetSize
		}
	}()

	return ch, nil
}

func showUploadFile(name string, uploadPath string, size int64) error {
	ch, err := uploadFile(name, uploadPath)
	if err != nil {
		return err
	}
	barHeader := NewProcessBarHeader("Upload", pt.Base(name))
	bar := NewProcessBar(barHeader, size)
	for {
		currentSize := <-ch
		if currentSize == -1 {
			break
		}
		bar.ShowProcess(currentSize)
	}
	return nil
}

func uploadDirAll(name string, uploadPath string) error {
	info, err := os.Stat(name)
	if err != nil {
		return err
	}
	err = createItem(info, uploadPath)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		err = showUploadFile(name, uploadPath, info.Size())
		return err
	}
	children, err := os.ReadDir(name)
	if err != nil {
		return err
	}
	for _, child := range children {
		err = uploadDirAll(name+"/"+child.Name(), uploadPath+"/"+pt.Base(name))
		if err != nil {
			return err
		}
	}
	return nil
}

func createItem(info os.FileInfo, uploadPath string) error {
	if IsItemExist(uploadPath + "/" + info.Name()) {
		return errors.New("\"" + info.Name() + "\" has already existed in current path. ")
	}
	var itemSize int64
	var itemType pb.ItemType
	if info.IsDir() {
		itemType = pb.ItemType_Directory
	} else {
		itemType = pb.ItemType_File
		itemSize = info.Size()
	}
	_, err := CNetDiskClient.CreateItem(context.Background(), &pb.CreateItemRequest{
		Info: &pb.ItemInfo{
			Name: info.Name(),
			Type: itemType,
			Size: itemSize,
			Path: uploadPath,
		},
	})
	return err
}

func UploadItem(name string, uploadPath string) error {
	stat, err := os.Stat(name)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		err = createItem(stat, uploadPath)
		if err != nil {
			return err
		}
		err = showUploadFile(name, uploadPath, stat.Size())
		if err != nil {
			return err
		}
	} else {
		err = uploadDirAll(name, uploadPath)
		if err != nil {
			return err
		}
	}
	return nil
}
