package server

import (
	"context"
	service "gochat_db/internal/service"
	pb "gochat_db/proto"
	"log"
)

type GochatByRpcDBServer struct{}

func (s *GochatByRpcDBServer) Login(ctx context.Context, req *pb.GochatDBLoginRequest) (*pb.GochatDBLoginResponse, error) {
	rsp := &pb.GochatDBLoginResponse{}
	err := service.Login(ctx, req, rsp)
	return rsp, err
}

func (s *GochatByRpcDBServer) Register(ctx context.Context, req *pb.GochatDBRegisterRequest) (*pb.GochatDBRegisterResponse, error) {
	log.Println("%v", req)
	rsp := &pb.GochatDBRegisterResponse{
		Errcode: 0,
		Errmsg:  "",
	}
	err := service.Register(ctx, req, rsp)
	return rsp, err
}
