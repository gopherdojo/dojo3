# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -coverprofile=profile
GOCOVER=$(GOCMD) tool cover -html=profile
GOGET=$(GOCMD) get
BINARY_NAME=pget
BINARY_UNIX=$(BINARY_NAME)_unix
 build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
	$(GOCOVER)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
