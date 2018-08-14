GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODOCCMD=godoc
GODOCPORT=6060
BINARY_NAME=imgconv

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f `git ls-files --others testdata/`
	git checkout HEAD testdata
doc:
	(sleep 1; open http://localhost:$(GODOCPORT)/pkg/github.com/gopherdojo/dojo3) &
	$(GODOCCMD) -http ":$(GODOCPORT)"
