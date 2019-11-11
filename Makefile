.PHONY: all build deploy

all: build

build:
	go generate
	go build ./cmd/kilabit

deploy: build
	rsync ./kilabit gcp-webserver:~/bin/kilabit

deploy-local: build
	rsync ./kilabit $(GOBIN)/
