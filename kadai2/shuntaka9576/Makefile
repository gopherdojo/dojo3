GOCMD=go
GORUN=go run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GODOCCMD=godoc
GODOCPORT=3000
BINARY_NAME=imageConverter
BINARY_WIN=imageConverter.exe

all:clean test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
	$(GOBUILD) -o $(BINARY_WIN) -v
run-topng:
	$(GORUN) main.go -f jpg -t png
run-tojpg:
	$(GORUN) main.go -f png -t jpg
test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_WIN)
doc:
	$(GODOCCMD) -http ":$(GODOCPORT)"
coverage-all:
	coverage-cli coverage-converter
coverage-cli:
	$(GOTEST) -coverprofile=cover_cli.out ./cli
	$(GOTOOL) cover -html=cover_cli.out -o cover_cli.html
coverage-converter:
	$(GOTEST) -coverprofile=cover_converter.out ./converter
	$(GOTOOL) cover -html=cover_converter.out -o cover_converter.html

