package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "gochat_db/proto"
	"gochat_db/server"
)

func main() {
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}
	log.Println("net listen...")
	grpcServer := grpc.NewServer()

	pb.RegisterGochatByRpcDBServer(grpcServer, &server.GochatByRpcDBServer{})

	if err = grpcServer.Serve(listener);err != nil {
		log.Fatalf("grpc server err: %v", err)
	}
}