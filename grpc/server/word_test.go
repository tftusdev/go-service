package server

import (
	"context"
	word_pb "go-service/grpc/pb/word"
	"go-service/service"
	"reflect"
	"testing"
)

func Test_wordServer_GetTop10MostUsedWords(t *testing.T) {
	type fields struct {
		Service                 service.WordService
		UnimplementedWordServer word_pb.UnimplementedWordServer
	}
	type args struct {
		ctx     context.Context
		request *word_pb.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *word_pb.Response
		wantErr bool
	}{
		{
			"Success",
			fields{
				Service: service.NewWordService(),
			},
			args{
				ctx: context.TODO(),
				request: &word_pb.Request{
					Text: "a bb bb ccc ccc ccc",
				},
			},
			&word_pb.Response{
				WordCountList: []*word_pb.WordCount{
					{
						Word:  "ccc",
						Count: 3,
					},
					{
						Word:  "bb",
						Count: 2,
					},
					{
						Word:  "a",
						Count: 1,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewWordServer(
				tt.fields.Service,
			)
			got, err := s.GetTop10MostUsedWords(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("wordServer.GetTop10MostUsedWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordServer.GetTop10MostUsedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
