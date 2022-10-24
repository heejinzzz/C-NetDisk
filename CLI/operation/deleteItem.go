package operation

import (
	pb "C-NetDisk/CLI/proto"
	"context"
	"errors"
)

func DeleteItem(name string) error {
	if !IsItemExist(name) {
		return errors.New("\"" + name + "\" doesn't exist. ")
	}
	_, err := CNetDiskClient.DeleteItem(context.Background(), &pb.DeleteItemRequest{
		Name: name,
	})
	return err
}
