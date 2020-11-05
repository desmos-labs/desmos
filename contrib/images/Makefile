all: desmosd-env

desmosd-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmosd-env desmosd-env

.PHONY: all desmosd-env
