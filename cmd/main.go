package main

import (
	"fmt"
	"go-service/api"
	"go-service/config"
	"go-service/grpc"
	"log"
	"net"
)

func main() {

	cfg, err := config.ReadConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to read env: %v", err)
	}

	var errs chan error

	r := api.Setup()

	go func() {
		errs <- r.Run(fmt.Sprintf(":%d", cfg.RestPort))
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.Setup()

	go func() {
		fmt.Printf("\ngRPC server listening at %v\n", lis.Addr())
		errs <- s.Serve(lis)
	}()

	if err = <-errs; err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	s.GracefulStop()
}
