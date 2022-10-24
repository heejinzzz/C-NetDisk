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

func DownloadItem(name string, savePath string) error {
	if !IsItemExist(name) {
		return errors.New("\"" + name + "\" doesn't exist")
	}
	err := DownloadDirAll(name, savePath)
	return err
}

func DownloadDirAll(name string, savePath string) error {
	info, err := GetItemInfo(name)
	if err != nil {
		return err
	}
	if info.Type == pb.ItemType_File {
		err = ShowDownloadFile(name, info.Size, savePath)
		if err != nil {
			return err
		}
	} else {
		err = os.MkdirAll(savePath+"/"+pt.Base(name), 0777)
		if err != nil {
			return err
		}
		for _, child := range info.Children {
			err = DownloadDirAll(child.Name, savePath+"/"+pt.Base(name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ShowDownloadFile(name string, size int64, savePath string) error {
	ch, err := downloadFile(name, savePath)
	if err != nil {
		return err
	}
	barHeader := NewProcessBarHeader("Download", name)
	bar := NewProcessBar(barHeader, size)
	for downloadSize := range ch {
		bar.ShowProcess(downloadSize)
	}
	return nil
}

func downloadFile(name string, savePath string) (chan int64, error) {
	stream, err := CNetDiskClient.DownloadFile(context.Background(), &pb.DownloadFileRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(savePath+"/"+pt.Base(name), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	ch := make(chan int64, DeliverSizeChannelLength)

	// receive file data
	go func() {
		defer func() {
			err := file.Close()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(2)
			}
		}()
		downloadSize := int64(0)
		writer := bufio.NewWriter(file)
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(ch)
				return
			}
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
			n, err := writer.Write(res.Data)
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(2)
			}
			downloadSize += int64(n)
			ch <- downloadSize
		}
	}()

	return ch, nil
}
