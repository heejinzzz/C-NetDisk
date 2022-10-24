package operation

import (
	pb "C-NetDisk/CLI/proto"
	"context"
)

func UserRegister(username string, password string) error {
	_, err := CNetDiskClient.UserRegister(context.Background(), &pb.UserRegisterRequest{
		Name:     username,
		Password: password,
	})
	return err
}
