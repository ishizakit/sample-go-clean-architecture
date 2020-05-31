# Const
#===============================================================
BIN_DIR  := deployments/bin
MAIN_DIR := app/main
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)
SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

include .make/*.mk

.PHONY: dep
dep:
	go get -u golang.org/x/lint/golint

.PHONY: lint
lint:
	golint -set_exit_status $$(go list ./...)
	go vet ./...

# MAIN_PATHを指定する
.PHONY: build
build:
	$(eval BIN_PATH := $(subst $(MAIN_DIR)/,$(BIN_DIR)/,$(subst /main.go,,$(MAIN_PATH))))
	@mkdir -p $(dir $(BIN_PATH));\
	rm -f $(BIN_PATH);\
	echo Build to $(BIN_PATH);\
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o $(BIN_PATH) $(MAIN_PATH)

.PHONY: test
test:
	@go test -v ./...
