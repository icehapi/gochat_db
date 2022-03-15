package service

import (
	"context"
	pb "gochat_db/proto"
	"gochat_db/utils"
	"log"
)

func Login(ctx context.Context, req *pb.GochatDBLoginRequest, rsp *pb.GochatDBLoginResponse) error {
	return nil
}

func Register(ctx context.Context, req *pb.GochatDBRegisterRequest, rsp *pb.GochatDBRegisterResponse) error {
	rsp.Authtoken = utils.GetRandomToken(32)
	log.Println("request body:", req)
	return nil
}
