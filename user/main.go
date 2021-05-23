package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"user/common"
	"user/handler"
	"user/internal/repository"
	s "user/internal/service"
	pb "user/proto"
)

func main() {

	// Tracing
	t, io, err := common.NewTracer("localhost:6831", "go-shop.user")
	if err != nil {
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Create service
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8081"),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	srv.Init()

	userService := s.NewUserService(repository.NewUserRepository())
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{UserService: userService})

	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
