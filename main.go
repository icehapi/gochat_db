package main

import (
	"gochat_db/internal/server"
	pb "gochat_db/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}
	log.Println("net listen...")
	grpcServer := grpc.NewServer()

	pb.RegisterGochatDBServer(grpcServer, &server.GochatByRpcDBServer{})

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("grpc server err: %v", err)
	}
}
