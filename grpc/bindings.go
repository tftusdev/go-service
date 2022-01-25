package grpc

import (
	word_pb "go-service/grpc/pb/word"
	"go-service/grpc/server"
	"go-service/service"
)

func bindWord() word_pb.WordServer {
	service := service.NewWordService()
	return server.NewWordServer(service)
}
