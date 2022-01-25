test:
	go test ./... -v -cover

build:
	docker build --build-arg PORT=$(PORT) -t go-service .
	docker run --env-file .env --network="host" -d --name go-service go-service

run:
	go run cmd/main.go
