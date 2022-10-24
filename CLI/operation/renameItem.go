package operation

import (
	pb "C-NetDisk/CLI/proto"
	"context"
	"errors"
)

func RenameItem(name string, newItemName string) error {
	if !IsItemExist(name) {
		return errors.New("\"" + name + "\" doesn't exist. ")
	}
	_, err := CNetDiskClient.RenameItem(context.Background(), &pb.RenameItemRequest{
		Name:        name,
		NewFileName: newItemName,
	})
	return err
}
