package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"purchase/internal/server"
	"purchase/internal/service"
	pb "purchase/proto"
)

func main()  {
	// Create service
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8081"),
	)

	srv.Init()

	userService := service.New()
	err := pb.RegisterPurchaseHandler(srv.Server(), &server.PurchaseServer{PurchaseService: userService})

	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
