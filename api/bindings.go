package api

import (
	"go-service/api/handler"
	"go-service/service"
)

func bindWord() handler.WordHandler {
	service := service.NewWordService()
	return handler.NewWordHandler(service)
}
