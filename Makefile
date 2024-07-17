IMAGE_NAME := osag1e.dev
CONTAINER_NAME := osag1e.dev-container

lint:
	@golangci-lint run ./...

docker:
	@docker build -t $(IMAGE_NAME) .
	@docker run -it --rm -p 8080:8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

clean: 
	@rm -rf bin

