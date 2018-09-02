# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=ozu

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v -cover ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)