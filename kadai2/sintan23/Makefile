NAME    := convertImage
VERSION := 0.0.1

SRC_PATH   := ./src
SRC        := $(SRC_PATH)/*.go
SRC_TEST   := $(SRC_PATH)/convert
BIN_PATH   := ./bin
OUTPUT_BIN := $(BIN_PATH)/$(NAME)

.PHONY: $(/bin/bash egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')


all: test build

.PHONY: test
test:
	## test
	go test -v ${SRC_TEST} -coverprofile cover.out
	go tool cover -func cover.out

.PHONY: build
build:
	## build
	mkdir -p $(BIN_PATH)
	go build -o $(OUTPUT_BIN) $(SRC)
	chmod 755 $(OUTPUT_BIN)

.PHONY: run
run:
	## run
	rm -f ./_data/jpg*.png
	go run $(SRC)

.PHONY: clean
clean:
	## clean
	rm -rf $(BIN_PATH)/*
