all: desmos-env relayer

desmos-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmos-env desmos-env

relayer:
	docker build --tag relayer relayer

.PHONY: all desmos-env relayer
