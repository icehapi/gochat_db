package server

import (
	"context"
	// "log"
	service "gochat_db/service"
	pb "gochat_db/proto"
)

type GochatByRpcDBServer struct {}

func (s *GochatByRpcDBServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	rsp := &pb.LoginResponse{}
	err := service.Login(ctx, req, rsp)
	return rsp, err
}

func (s *GochatByRpcDBServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	rsp := &pb.RegisterResponse{}
	err := service.Register(ctx, req, rsp)
	return rsp, err
}


