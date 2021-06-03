# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) mod download
BINARY_NAME=minerva
BINARY_PATH=bin/
BINARY_LINUX=$(BINARY_NAME)_linux
ENTRY_POINT=cmd/cli/minerva.go


all: clean test build
build:
		$(GOBUILD) -o $(BINARY_PATH)$(BINARY_NAME) -v $(ENTRY_POINT)
install:
		$(GOINSTALL) $(ENTRY_POINT)
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -r $(BINARY_PATH)$(BINARY_NAME)
run:
		$(GORUN) $(ENTRY_POINT)
deps:
		$(GOGET)


# Cross compilation
build-linux:
		echo "Not Implemented"
		# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
		echo "Not Implemented"
		# docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
