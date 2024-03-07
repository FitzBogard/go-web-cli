SHELL := /bin/bash

APP := go-web-cli

VERSION := $(shell git describe --tags --always)
FULL_COMMIT := $(shell git rev-parse HEAD)
RELEASE_DATE := $(shell git show -s --format=%cI)
BUILD_DATE = $(shell date '+%FT%T%z')
LDFLAGS = -ldflags "-X $(APP)/internal/pkg/version.Version=$(VERSION) -X $(APP)/internal/pkg/version.FullCommit=$(FULL_COMMIT) -X $(APP)/internal/pkg/version.ReleaseDate=$(RELEASE_DATE) -X $(APP)/internal/pkg/version.BuildDate=$(BUILD_DATE)"
IDL_DIR = ../api-idl/fee/

run:
	go run $(LDFLAGS) -race .

gen:
	go generate ./...

grpc_gen:
	kitex -module fx-engine-proxy -type protobuf -I $(IDL_DIR) fee.proto

release: clean
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ./bin/$(APP)-$(VERSION)

clean:
	rm -rf ./bin

lint:
	golangci-lint run ./... -v