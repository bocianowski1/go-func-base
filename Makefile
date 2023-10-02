build: 
	@go build -o main

run: build
	@docker-compose up -d
	@func start

test:
	@go test v ./...