#

A small service that accepts as input a body of text, such as that from a book, and returns the top ten most-used words along with how many times they occur in the text.

Rest and gRPC servers are implemented

## Requests

gRPCurl : 

grpcurl -plaintext -d '{"Text" : "a a bb ccc"}' localhost:3001 word.Word/GetTop10MostUsedWords

cURL :

curl --location --request POST 'localhost:3000/top10' \
--header 'Content-Type: text/plain' \
--data-raw 'a a bb ccc'


## Proto generation

protoc --go_out=. --go-grpc_out=. grpc/proto/word.proto