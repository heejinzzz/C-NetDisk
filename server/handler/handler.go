package handler

import (
	pb "C-NetDisk/server/proto"
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"os"
	pt "path"
	"sync"
)

func NewCNetDiskServer() *CNetDiskServer {
	return &CNetDiskServer{}
}

type CNetDiskServer struct{}

func (*CNetDiskServer) UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	log.Println("[INFO]Get CNetDisk.UserRegister Request: ", req)
	if isUserExist(req.Name) {
		return nil, errors.New("username has been used! Please change your username and retry. ")
	}
	if !isUsernameValid(req.Name) {
		return nil, errors.New("username is invalid! Please make sure your password is within the range of 4 to 15 in length and does not contain special characters. ")
	}
	err := createNewUser(req.Name, req.Password)
	if err != nil {
		return nil, errors.New("register failed. Error: " + err.Error())
	}
	return &pb.UserRegisterResponse{Msg: "register succeed!"}, nil
}

func (*CNetDiskServer) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	log.Println("[INFO]Get CNetDisk.UserLogin Request: ", req)
	if checkUserPassword(req.Name, req.Password) {
		return &pb.UserLoginResponse{Msg: "login succeed!"}, nil
	}
	return nil, errors.New("login failed. Username or Password is wrong. ")
}

func (*CNetDiskServer) GetItemInfo(ctx context.Context, req *pb.GetItemInfoRequest) (*pb.GetItemInfoResponse, error) {
	log.Println("[INFO]Get CNetDisk.GetItemInfo Request: ", req)
	itemInfo, err := getItemInfoFromMongo(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.GetItemInfoResponse{Info: itemInfo}, nil
}

func (*CNetDiskServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	log.Println("[INFO]Get CNetDisk.CreateItem Request: ", req)
	itemInfo, err := createItemInMongo(req.Info)
	if err != nil {
		return nil, errors.New("Create " + req.Info.Name + " in mongodb failed. Error: " + err.Error())
	}
	err = createItemInDisk(itemInfo)
	if err != nil {
		return nil, errors.New("Create file " + itemInfo.Name + " failed. Error: " + err.Error())
	}
	return &pb.CreateItemResponse{Msg: "CreateItem Succeed!"}, nil
}

func (*CNetDiskServer) UploadFile(stream pb.CNetDisk_UploadFileServer) error {
	ch := make(chan int64, 100)
	var wg sync.WaitGroup
	wg.Add(2)

	// receive file
	go func() {
		defer wg.Done()
		var file *os.File
		currentSize := int64(0)
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				ch <- -1
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			if file == nil {
				path := getItemPath(req.Name)
				file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0777)
				if err != nil {
					log.Fatalln(err)
				}
			}
			_, err = file.Write(req.Data)
			if err != nil {
				log.Fatalln(err)
			}
			currentSize += int64(len(req.Data))
			ch <- currentSize
		}
		if file != nil {
			err := file.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()

	// response to client
	go func() {
		defer wg.Done()
		for {
			currentSize := <-ch
			if currentSize == -1 {
				err := stream.Send(&pb.UploadFileResponse{ResSignal: pb.Signal_FinishReceive})
				if err != nil {
					log.Fatalln(err)
				}
				break
			}
			err := stream.Send(&pb.UploadFileResponse{ResSignal: pb.Signal_Receiving, GetSize: currentSize})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()

	wg.Wait()
	return nil
}

func (*CNetDiskServer) DownloadFile(req *pb.DownloadFileRequest, stream pb.CNetDisk_DownloadFileServer) error {
	log.Println("[INFO]Get CNetDisk.DownloadFile Request: ", req)
	path := getItemPath(req.Name)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	reader := bufio.NewReader(file)
	for {
		buf := make([]byte, ReadBytesEachTime)
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = stream.Send(&pb.DownloadFileResponse{Data: buf[:n]})
		if err != nil {
			return err
		}
	}
	return nil
}

func (*CNetDiskServer) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	log.Println("[INFO]Get CNetDisk.DeleteItem Request: ", req)
	err := deleteItem(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteItemResponse{Msg: "Delete " + req.Name + " Succeed!"}, nil
}

func (*CNetDiskServer) RenameItem(ctx context.Context, req *pb.RenameItemRequest) (*pb.RenameItemResponse, error) {
	log.Println("[INFO]Get CNetDisk.RenameItem Request: ", req)
	err := renameItem(req.Name, req.NewFileName)
	if err != nil {
		return nil, err
	}
	return &pb.RenameItemResponse{Msg: "Rename " + pt.Base(req.Name) + " to " + req.NewFileName + " Succeed!"}, nil
}
