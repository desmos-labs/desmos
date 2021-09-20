desmos-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmos-env desmos-env

desmos-build:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmos-build desmos-build

desmos-python:
	docker build --tag desmoslabs/desmos-python desmos-python

desmos-cosmovisor:
	docker build \
		--build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) \
		--build-arg DESMOS_VERSION=$(DESMOS_VERSION) \
		--tag desmoslabs/desmos-cosmovisor:$(DESMOS_VERSION) \
		desmos-cosmovisor

all: desmos-env desmos-python desmos-cosmovisor desmos-build

.PHONY: all desmos-env desmos-python desmos-cosmovisor desmos-build
