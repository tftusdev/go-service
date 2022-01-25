package handler

import (
	"go-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordHandler interface {
	GetTop10MostUsedWords(c *gin.Context)
}

type wordHandler struct {
	Service service.WordService
}

func NewWordHandler(service service.WordService) WordHandler {
	return &wordHandler{
		Service: service,
	}
}

func (h *wordHandler) GetTop10MostUsedWords(c *gin.Context) {

	bytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := h.Service.GetTop10MostUsedWords(string(bytes))

	c.JSON(http.StatusOK, response)
}
