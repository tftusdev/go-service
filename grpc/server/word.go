package server

import (
	"context"
	word_pb "go-service/grpc/pb/word"
	"go-service/service"
)

func NewWordServer(service service.WordService) *wordServer {
	return &wordServer{
		Service: service,
	}
}

type wordServer struct {
	Service service.WordService
	word_pb.UnimplementedWordServer
}

func (s *wordServer) GetTop10MostUsedWords(ctx context.Context, request *word_pb.Request) (*word_pb.Response, error) {

	wordCountList := s.Service.GetTop10MostUsedWords(request.Text)

	return &word_pb.Response{
		WordCountList: wordCountList,
	}, nil
}
