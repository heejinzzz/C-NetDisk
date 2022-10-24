package operation

import (
	pb "C-NetDisk/CLI/proto"
	"context"
	"log"
	pt "path"
)

func GetItemInfo(name string) (*pb.ItemInfo, error) {
	response, err := CNetDiskClient.GetItemInfo(context.Background(), &pb.GetItemInfoRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return response.Info, nil
}

func IsItemExist(name string) bool {
	currentPath := pt.Dir(name)
	info, err := GetItemInfo(currentPath)
	if err != nil {
		log.Fatalln(err)
	}
	children := info.Children
	for _, child := range children {
		if child.Name == name {
			return true
		}
	}
	return false
}
