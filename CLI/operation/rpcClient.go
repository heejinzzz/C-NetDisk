package operation

import (
	pb "C-NetDisk/CLI/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ServerAddr = "localhost:6732"

var CNetDiskClient pb.CNetDiskClient

func ConnectCNetDiskServer(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	CNetDiskClient = pb.NewCNetDiskClient(conn)
	return nil
}
