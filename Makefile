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
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmosd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmoscli
	# go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmoskeyutil

########################################
### Build

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmod.exe ./cmd/desmosd
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmoscli.exe ./cmd/desmoscli
else
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmosd ./cmd/desmosd
	go build -mod=readonly $(BUILD_FLAGS) -o ./build/desmoscli ./cmd/desmoscli
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/gaiad -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf snapcraft-local.yaml build/

distclean: clean
	rm -rf vendor/

########################################
### Testing

test: test-unit test-build
test-all: test test-race test-cover

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_NOSIMULATION) -tags='ledger test_ledger_mock'

test-build: build
	@go test -mod=readonly -p 4 `go list ./cli_test/...` -tags=cli_test -v

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' ./...

test-cover:
	@go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...


lint: golangci-lint
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs goimports -w -local github.com/desmos-labs/desmos

benchmark:
	@go test -mod=readonly -bench=. ./...

########################################
### Local validator nodes using docker and docker-compose

build-docker-desmosdnode:
	$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: build-linux localnet-stop
	@if ! [ -f build/node0/desmosd/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/desmosd:Z desmoslabs/desmos desmosd testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test ; fi
	docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down


# include simulations
include Makefile.simulations

.PHONY: all build-linux install \
	go-mod-cache clean build \
	test test-all test-cover test-unit test-race