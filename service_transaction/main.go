package main

import (
	"github.com/go-micro/plugins/v4/server/grpc"

	"service-transaction/handler"
	pb "service-transaction/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
	// "service-transaction/handler"
	// pb "service-transaction/proto"
	// "go-micro.dev/v4"
	// "go-micro.dev/v4/codec/grpc"
	// "go-micro.dev/v4/logger"
)

var (
	service = "service-transaction"
	version = "latest"
)

func main() {
	server := grpc.NewServer((server.Address(":9000")))

	src := micro.NewService(
		micro.Server(server),
	)

	src.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	err := pb.RegisterServiceTransactionHandler(src.Server(), new(handler.ServiceTransaction))
	if err != nil {
		logger.Fatal(err)
	}

	err = src.Run()
	if err != nil {
		logger.Fatal(err)
	}
	// Create service
	// srv := micro.NewService(
	// )
	// srv.Init(
	// 	micro.Name(service),
	// 	micro.Version(version),
	// )

	// // Register handler
	// if err := pb.RegisterServiceTransactionHandler(srv.Server(), new(handler.ServiceTransaction)); err != nil {
	// 	logger.Fatal(err)
	// }
	// // Run service
	// if err := srv.Run(); err != nil {
	// 	logger.Fatal(err)
	// }
}
