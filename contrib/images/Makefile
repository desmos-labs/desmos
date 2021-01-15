all: desmos-env

desmos-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmos-env desmos-env

.PHONY: all desmos-env
