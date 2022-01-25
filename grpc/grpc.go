package grpc

import (
	word_pb "go-service/grpc/pb/word"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Setup() *grpc.Server {

	s := grpc.NewServer()

	word_pb.RegisterWordServer(s, bindWord())

	reflection.Register(s)

	return s
}
