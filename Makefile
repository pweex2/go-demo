# Go settings
BINARY_NAME=go-demo
BUILD_DIR=bin

# Targets
.PHONY: all test build clean docker docker-push deploy-local

all: build

test:
	go test ./...

# Build for current platform (Mac/Windows)
build:
ifeq ($(shell uname -s),Darwin)
	go build -v -o $(BUILD_DIR)/$(BINARY_NAME)
else
	GOOS=linux GOARCH=amd64 go build -v -o $(BUILD_DIR)/$(BINARY_NAME)
endif

build-mac:
	GOOS=darwin GOARCH=amd64 go build -v -o $(BUILD_DIR)/$(BINARY_NAME)

build-linux:
	GOOS=linux GOARCH=amd64 go build -v -o $(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)

docker:
	docker build -t puahvvx2/cicd-go-demo:latest .

docker-push: docker
	docker push puahvvx2/cicd-go-demo:latest

deploy-local:
	@echo "🔧 Running deploy.sh..."
	@test -f $(BUILD_DIR)/$(BINARY_NAME) || (echo "❌ Binary not found. Run make build first."; exit 1)
	@if [ ! -x ./deploy.sh ]; then chmod +x ./deploy.sh; fi
	./deploy.sh

deploy-remote:
	@echo "🚀 Simulating remote deployment..."
	@echo "scp -i ~/.ssh/id_rsa bin/go-demo-linux user@remote:/usr/local/bin/go-demo"
	@echo "ssh -i ~/.ssh/id_rsa user@remote 'pkill go-demo || true && nohup /usr/local/bin/go-demo &'"
	@echo "rsync -avz -e 'ssh -i ~/.ssh/id_rsa' bin/go-demo-linux user@remote:/usr/local/bin/go-demo"


lint:
	golangci-lint run ./...

# Makefile
test:
	go test ./...