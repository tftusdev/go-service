package api

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	wordHandler := bindWord()
	r.POST("/top10", wordHandler.GetTop10MostUsedWords)

	return r
}
