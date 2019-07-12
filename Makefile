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

ldflags = github.com/cosmos/cosmos-sdk/version.Name=desmos \
 	-X github.com/cosmos/cosmos-sdk/version.ServerName=desmosd \
 	-X github.com/cosmos/cosmos-sdk/version.ClientName=desmoscli \
 	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
  	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags)"

all: lint install

install: go.sum
		GO111MODULE=on go install $(BUILD_FLAGS) ./cmd/desmosd
		GO111MODULE=on go install $(BUILD_FLAGS) ./cmd/desmoscli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify
