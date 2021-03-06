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
build-all:
		echo "Not Implemented"
		# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
		compile:
		# echo "Compiling for every OS and Platform"
		# GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
		# GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
		# GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
docker-build:
		echo "Not Implemented"
		# docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
