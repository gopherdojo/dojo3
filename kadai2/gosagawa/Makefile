VERSION=1.0.0
TARGETS_NOVENDOR=$(shell go list ./... | grep -v /vendor/)

all: bin/imgconv

bin/imgconv: main.go
	go build -o bin/imgconv main.go

fmt:
	@echo ./... | xargs go fmt

check:
	go test -v -cover ./...

clean:
	rm -rf bin/*

