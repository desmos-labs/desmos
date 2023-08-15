#!/usr/bin/make -f

PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --always) | sed 's/^v//')
TENDERMINT_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
SIMAPP = ./app
MOCKS_DIR = $(CURDIR)/tests/mocks
HTTPS_GIT := https://github.com/desmos-labs/desmos.git
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
BENCH_COUNT ?= 5
REF_NAME ?= $(shell git symbolic-ref HEAD --short | tr / - 2>/dev/null)

export GO111MODULE = on

###############################################################################
###                                Build flags                              ###
###############################################################################

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

# These lines here are essential to include the muslc library for static linking of libraries
# (which is needed for the wasmvm one) available during the build. Without them, the build will fail.
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# Process linker flags
ldflags = -X 'github.com/cosmos/cosmos-sdk/version.Name=Desmos' \
 	-X 'github.com/cosmos/cosmos-sdk/version.AppName=desmos' \
 	-X 'github.com/cosmos/cosmos-sdk/version.Version=$(VERSION)' \
    -X 'github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)' \
  	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
  	-X "github.com/cometbft/cometbft/version.TMCoreSemVer=$(TENDERMINT_VERSION)"

ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

ifneq ($(GOSUM),)
  ldflags += -X github.com/cosmos/cosmos-sdk/version.VendorDirHash=$(shell $(GOSUM) go.sum)
endif

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X 'github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb'
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X 'github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb'
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

###############################################################################
###                                   All                                   ###
###############################################################################

all: tools build lint test

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

create-builder: go.sum
	$(MAKE) -C contrib/images desmos-builder CONTEXT=$(CURDIR)

build-alpine: create-builder
	mkdir -p $(BUILDDIR)
	$(DOCKER) build -f Dockerfile --rm --tag desmoslabs/desmos-alpine .
	$(DOCKER) create --name desmos-alpine --rm desmoslabs/desmos-alpine
	$(DOCKER) cp desmos-alpine:/usr/bin/desmos $(BUILDDIR)/desmos
	$(DOCKER) rm desmos-alpine

build-linux: create-builder
	mkdir -p $(BUILDDIR)
	$(DOCKER) build -f Dockerfile-ubuntu --rm --tag desmoslabs/desmos-linux .
	$(DOCKER) create --name desmos-linux desmoslabs/desmos-linux
	$(DOCKER) cp desmos-linux:/usr/bin/desmos $(BUILDDIR)/desmos
	$(DOCKER) rm desmos-linux

build-reproducible: go.sum
	$(DOCKER) rm latest-build || true
	$(DOCKER) run --volume=$(CURDIR):/sources:ro \
        --env TARGET_PLATFORMS='linux/amd64 linux/arm64 darwin/amd64 windows/amd64' \
        --env APP=desmos \
        --env VERSION=$(VERSION) \
        --env COMMIT=$(COMMIT) \
        --env LEDGER_ENABLED=$(LEDGER_ENABLED) \
        --name latest-build cosmossdk/rbuilder:latest
	$(DOCKER) cp -a latest-build:/home/builder/artifacts/ $(CURDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

mockgen:
	@go install github.com/golang/mock/mockgen@v1.6.0
	@./scripts/mockgen.sh
.PHONY: mocks

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@go mod tidy

clean:
	rm -rf \
	$(BUILDDIR)/ \
	artifacts/ \
	tmp-swagger-gen/

distclean: clean tools-clean

.PHONY: distclean clean

###############################################################################
###                              Documentation                              ###
###############################################################################

update-swagger-docs: statik
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
		echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
		exit 1;\
	else \
		echo "\033[92mSwagger docs are in sync\033[0m";\
	fi
.PHONY: update-swagger-docs

build-docs:
	@cd docs && DOCS_DOMAIN=docs.desmos.network bash ./build-all.sh

.PHONY: build-docs

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test: test-unit
test-all: test-unit test-ledger-mock test-race test-cover

TEST_PACKAGES=./...
TEST_TARGETS := test-unit test-unit-amino test-unit-proto test-ledger-mock test-race test-ledger test-race

# Test runs-specific rules. To add a new test target, just add
# a new rule, customise ARGS or TEST_PACKAGES ad libitum, and
# append the new rule to the TEST_TARGETS list.
test-unit: test_tags += cgo ledger test_ledger_mock norace
test-unit-amino: test_tags += ledger test_ledger_mock test_amino norace
test-ledger: test_tags += cgo ledger norace
test-ledger-mock: test_tags += ledger test_ledger_mock norace
test-race: test_tags += cgo ledger test_ledger_mock
test-race: ARGS=-race
test-race: TEST_PACKAGES=$(PACKAGES_NOSIMULATION)
$(TEST_TARGETS): run-tests

ARGS += -tags "$(test_tags)"
SUB_MODULES = $(shell find . -type f -name 'go.mod' -print0 | xargs -0 -n1 dirname | sort)
CURRENT_DIR = $(shell pwd)

run-tests:
ifneq (,$(shell which tparse 2>/dev/null))
	go test -mod=readonly -json $(ARGS) $(TEST_PACKAGES) | tparse
else
	go test -mod=readonly $(ARGS) $(TEST_PACKAGES)
endif

.PHONY: run-tests test test-all $(TEST_TARGETS)

test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@cd ${CURRENT_DIR}/app && go test -mod=readonly -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h

test-sim-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.desmos/config/genesis.json will be used."
	@cd ${CURRENT_DIR}/app && go test -mod=readonly -run TestFullAppSimulation -Genesis=${HOME}/.desmos/config/genesis.json \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@cd ${CURRENT_DIR}/app && $(BINDIR)/runsim -Jobs=4 -SimAppPkg=. -ExitOnFail 50 5 TestAppImportExport

test-sim-after-import: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	@cd ${CURRENT_DIR}/app && $(BINDIR)/runsim -Jobs=4 -SimAppPkg=. -ExitOnFail 50 5 TestAppSimulationAfterImport

test-sim-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.desmos/config/genesis.json will be used."
	@cd ${CURRENT_DIR}/app && $(BINDIR)/runsim -Genesis=${HOME}/.desmos/config/genesis.json -SimAppPkg=. -ExitOnFail 400 5 TestFullAppSimulation

test-sim-multi-seed-long: runsim
	@echo "Running long multi-seed application simulation. This may take awhile!"
	@cd ${CURRENT_DIR}/app && $(BINDIR)/runsim -Jobs=4 -SimAppPkg=. -ExitOnFail 500 50 TestFullAppSimulation

test-sim-multi-seed-short: runsim
	@echo "Running short multi-seed application simulation. This may take awhile!"
	@cd ${CURRENT_DIR}/app && $(BINDIR)/runsim -Jobs=4 -SimAppPkg=. -ExitOnFail 50 10 TestFullAppSimulation

test-sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@cd ${CURRENT_DIR}/app && @go test -mod=readonly -benchmem -bench=BenchmarkInvariants -run=^$ \
	-Enabled=true -NumBlocks=1000 -BlockSize=200 \
	-Period=1 -Commit=true -Seed=57 -v -timeout 24h

.PHONY: \
test-sim-nondeterminism \
test-sim-custom-genesis-fast \
test-sim-import-export \
test-sim-after-import \
test-sim-custom-genesis-multi-seed \
test-sim-multi-seed-short \
test-sim-multi-seed-long \
test-sim-benchmark-invariants

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_COMMIT ?= true

test-sim-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkFullAppSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h

test-sim-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkFullAppSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out

.PHONY: test-sim-profile test-sim-benchmark

test-cover:
	@export VERSION=$(VERSION); bash -x contrib/test_cover.sh
.PHONY: test-cover

becnchstat_cmd=golang.org/x/perf/cmd/benchstat
benchmark:
	@go test -mod=readonly -bench=. -count=$(BENCH_COUNT) -run=^a  ./... > bench-$(REF_NAME).txt
	@test -e bench-master.txt && go run $(becnchstat_cmd) bench-master.txt bench-$(REF_NAME).txt || go run $(becnchstat_cmd) bench-$(REF_NAME).txt
.PHONY: benchmark

test-mutation:
	@bash scripts/mutation-test.sh $(MODULES)
.PHONY: test-mutation

###############################################################################
###                                Linting                                  ###
###############################################################################
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	@echo "--> Running linter"
	@go run $(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go run $(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

.PHONY: lint lint-fix

format:
	find . -name '*.go' -type f -not -path "./client/docs/statik*" -not -path "*.git*" -not -name '*.pb.go' -not -name '*_mocks.go' -not -name '*.pulsar.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./client/docs/statik*" -not -path "*.git*" -not -name '*.pb.go' -not -name '*_mocks.go' -not -name '*.pulsar.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./client/docs/statik*" -not -path "*.git*" -not -name '*.pb.go' -not -name '*_mocks.go' -not -name '*.pulsar.go' | xargs goimports -w -local github.com/desmos-labs/desmos
.PHONY: format

###############################################################################
###                                  Types                                  ###
###############################################################################

EVMOS_URL 		 = https://raw.githubusercontent.com/evmos/ethermint/v0.17.0
CRYPTO_TYPES 	 = types/crypto

update-deps-types:
	@mkdir -p $(CRYPTO_TYPES)/ethsecp256k1
	@curl -sSL $(EVMOS_URL)/proto/ethermint/crypto/v1/ethsecp256k1/keys.proto > proto/ethermint/crypto/v1/ethsecp256k1/keys.proto
	@curl -sSL $(EVMOS_URL)/crypto/ethsecp256k1/ethsecp256k1.go > $(CRYPTO_TYPES)/ethsecp256k1/ethsecp256k1.go

.PHONY: update-deps-types

###############################################################################
###                                Protobuf                                 ###
###############################################################################
protoVer=0.11.6
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace --user $(shell id -u):$(shell id -g) $(protoImageName)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh
	$(MAKE) update-swagger-docs

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=master

proto-update-deps:
	$(DOCKER) run --rm -v $(CURDIR)/proto:/workspace --workdir /workspace $(protoImageName) buf mod update

.PHONY: proto-all proto-gen proto-lint proto-check-breaking proto-update-deps

###############################################################################
###                                Localnet                                 ###
###############################################################################

build-docker-desmosnode:
	$(MAKE) -C networks/local

# Setups 4 folders representing each one the genesis state of a testnet node
setup-localnet: build-linux
	if ! [ -f build/node0/desmos/config/genesis.json ]; then $(BUILDDIR)/desmos testnet \
		-o ./build --starting-ip-address 192.168.255.2 --keyring-backend=test \
		--v=$(if $(NODES),$(NODES),4) \
		--gentx-coin-denom=$(if $(COIN_DENOM),$(COIN_DENOM),"stake") \
		--minimum-gas-prices="0.000006$(if $(COIN_DENOM),$(COIN_DENOM),"stake")"; fi

# Starts a local 4-nodes testnet that should be used to test on-chain upgrades.
# It requires 3 arguments to work:
# 1. GENESIS_VERSION, which represents the Desmos version to be used when starting the testnet
# 2. GENESIS_URL, which represents the URL from where to download the testnet genesis status
# 3. UPGRADE_NAME, which represents the name of the upgrade to perform
upgrade-testnet-start: upgrade-testnet-stop
	$(CURDIR)/contrib/upgrade_testnet/start.sh 4 $(GENESIS_VERSION) $(GENESIS_URL) $(UPGRADE_NAME)

# Stops the 4-nodes testnet that should be used to test on-chain upgrades.
upgrade-testnet-stop:
	$(CURDIR)/contrib/upgrade_testnet/stop.sh

# Run a 4-node testnet locally
localnet-start: localnet-stop setup-localnet
	$(if $(shell docker inspect -f '{{ .Id }}' desmoslabs/desmos-env 2>/dev/null),$(info found image desmoslabs/desmos-env),$(MAKE) -C contrib/images desmos-env)
	docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down

.PHONY: all build-linux install \
	go-mod-cache clean build \
	test test-all test-cover test-unit test-race
