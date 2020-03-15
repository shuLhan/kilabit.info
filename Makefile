.PHONY: all build deploy

all: build

build:
	go generate
	go build ./cmd/www-kilabit

deploy: build
	rsync --progress ./www-kilabit aws-www:~/bin/www-kilabit

deploy-local: build
	rsync ./www-kilabit $(GOBIN)/

serve:
	DEBUG=1 go run ./cmd/www-kilabit
