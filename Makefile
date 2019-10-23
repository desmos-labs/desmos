PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

export GO111MODULE = on

include Makefile.ledger

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

# process linker flags

ldflags = -X "github.com/cosmos/cosmos-sdk/version.Name=Desmos" \
 	-X "github.com/cosmos/cosmos-sdk/version.ServerName=desmosd" \
 	-X "github.com/cosmos/cosmos-sdk/version.ClientName=desmoscli" \
 	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
  	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags)"

ifneq ($(GOSUM),)
ldflags += -X github.com/cosmos/cosmos-sdk/version.VendorDirHash=$(shell $(GOSUM) go.sum)
endif

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

all: lint install

########################################
### Install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmosd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmoscli
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/desmoskeyutil

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

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

########################################
### Testing

test: test_unit

test_unit:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_NOSIMULATION) -tags='ledger test_ledger_mock'