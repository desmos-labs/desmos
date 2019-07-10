include Makefile.ledger
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
