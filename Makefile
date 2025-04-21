# Go settings
BINARY_NAME=go-demo
BUILD_DIR=bin

# Targets
.PHONY: all test build clean docker docker-push

all: build

test:
	go test ./...

build:
	go build -v -o $(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)

docker:
	docker build -t puahvvx2/cicd-go-demo:latest .

docker-push: docker
	docker push puahvvx2/cicd-go-demo:latest