package handler

import (
	service "go-service/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_wordHandler_GetTop10MostUsedWords(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	wordHandler := NewWordHandler(service.NewWordService())
	r.POST("/top10", wordHandler.GetTop10MostUsedWords)

	type args struct {
		Text string
	}
	tests := []struct {
		name       string
		args       args
		StatusCode int
	}{
		{
			"Success",
			args{
				Text: "a a b",
			},
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/top10", strings.NewReader(tt.args.Text))
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)
			if resp.Code != tt.StatusCode {
				t.Errorf("wordServer.GetTop10MostUsedWords() = %v, want %v", resp.Code, tt.StatusCode)
			}
		})
	}
}
