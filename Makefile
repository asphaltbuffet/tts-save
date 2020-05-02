# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=tts-save
BINARY_UNIX=$(BINARY_NAME)_unix

all: t b build-linux
b:
	$(GOBUILD) -o $(BINARY_NAME) -v
t: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:



# Cross compilation
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
