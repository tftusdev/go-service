ARG REST_PORT=3000
ARG GRPC_PORT=3001

FROM golang:1.16-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE $REST_PORT
EXPOSE $GRPC_PORT

CMD ["./main"]
