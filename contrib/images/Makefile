desmos-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmos-env desmos-env

desmos-python:
	docker build --tag desmoslabs/desmos-python desmos-python

desmos-cosmovisor:
	docker build \
		--build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) \
		--build-arg DESMOS_VERSION=$(DESMOS_VERSION) \
		--tag desmoslabs/desmos-cosmovisor:$(DESMOS_VERSION) \
		desmos-cosmovisor

desmos-builder:
	docker build --tag desmoslabs/builder -f desmos-builder/Dockerfile $(CONTEXT)

all: desmos-env desmos-python desmos-cosmovisor desmos-builder

.PHONY: all desmos-env desmos-python desmos-cosmovisor desmos-builder
