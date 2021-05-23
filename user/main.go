package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"user/handler"
	"user/internal/repository"
	s "user/internal/service"
	pb "user/proto"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)

	srv.Init()

	userService := s.NewUserService(repository.NewUserRepository())
	err := pb.RegisterUserHandler(srv.Server(),&handler.User{UserService:userService})

	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
