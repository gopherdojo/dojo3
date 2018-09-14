GOCMD=go
GORUN=go run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GODOCCMD=godoc
GODOCPORT=3000
BINARY_NAME=typing
BINARY_WIN=typing.exe

all:clean test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
	$(GOBUILD) -o $(BINARY_WIN) -v
test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_WIN)
cov:
	$(GOTEST) ./... -race -coverprofile=coverage/c.out -covermode=atomic
	$(GOTOOL) cover -html=coverage/c.out -o coverage/index.html
	open coverage/index.html
doc:
	$(GODOCCMD) -http ":$(GODOCPORT)"