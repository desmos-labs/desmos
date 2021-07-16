all: desmos-python

desmos-python:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmos-python .

.PHONY: all desmos-python
