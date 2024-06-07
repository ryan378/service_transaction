package handler

import (
	"context"
	

	"go-micro.dev/v4/logger"

	pb "service-transaction/proto"
)

type ServiceTransaction struct{}

func (e *ServiceTransaction) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	logger.Infof("Received ServiceTransaction.Call request: %v", req)
	//rsp.Msg = "Hello " + req.Name
	rsp.Data = []*pb.LoginData{
		{
			Token: "",
		},
	}

	rsp.Message = req.Username
	return nil
}