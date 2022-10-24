package main

import (
	"C-NetDisk/server/handler"
	pb "C-NetDisk/server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const serverAddr = "0.0.0.0:6732"

func main() {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCNetDiskServer(grpcServer, handler.NewCNetDiskServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
