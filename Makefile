build-app:
	@go build -o bin/app ./cmd/web/

run: build-app
	@./bin/app

lint:
	@golangci-lint run ./...

clean: 
	@rm -rf bin

