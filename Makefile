build: 
	@go build -o main

run: build
	@docker-compose up -d
	@func start

test:
	@go test v ./...

clean:
	@docker-compose down
	@rm -rf main

deploy: build
	@bash ./deploy-func.sh