GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GODOCCMD=godoc
GODOCPORT=6060
BINARY_NAME=omikuji-server

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) ./...
cov:
	$(GOTEST) ./... -race -coverprofile=coverage/c.out -covermode=atomic
	$(GOTOOL) cover -html=coverage/c.out -o coverage/index.html
	open coverage/index.html
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
doc:
	(sleep 1; open http://localhost:$(GODOCPORT)/pkg/github.com/gopherdojo/dojo3) &
	$(GODOCCMD) -http ":$(GODOCPORT)"
