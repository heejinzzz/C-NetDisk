package operation

import (
	pb "C-NetDisk/CLI/proto"
	"context"
)

func UserLogin(username string, password string) error {
	_, err := CNetDiskClient.UserLogin(context.Background(), &pb.UserLoginRequest{
		Name:     username,
		Password: password,
	})
	return err
}
