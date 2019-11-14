#!/usr/bin/make -f

PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin

export GO111MODULE = on

include Makefile.ledger
include contrib/devtools/Makefile

########################################
### Build flags

ifneq ($(GOSUM),)
  ldflags += -X github.com/cosmos/cosmos-sdk/version.VendorDirHash=$(shell $(GOSUM) go.sum)
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

# Process linker flags
ldflags = -X "github.com/cosmos/cosmos-sdk/version.Name=Desmos" \
 	-X "github.com/cosmos/cosmos-sdk/version.ServerName=desmosd" \
 	-X "github.com/cosmos/cosmos-sdk/version.ClientName=desmoscli" \
 	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
  	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags)"

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

########################################
### All

all: lint install

########################################
### Install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmosd/main.go
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmoscli/main.go
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmoskeyutil/main.go

########################################
### Build

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmod.exe ./cmd/desmosd/main.go
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmoscli.exe ./cmd/desmoscli/main.go
else
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmosd ./cmd/desmosd/main.go
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmoscli ./cmd/desmoscli/main.go
endif

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
.PHONY: go-mod-cache

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@go mod tidy

lint: golangci-lint
	$(BINDIR)/golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify
.PHONY: lint

########################################
### Testing

test: test_unit

test_unit:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_NOSIMULATION) -tags='ledger test_ledger_mock'