.PHONY: all build deploy

all: build

build:
	go generate
	go build ./cmd/www-kilabit

www-kilabit-linux-amd64:
	go generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o $@ ./cmd/www-kilabit/

deploy: www-kilabit-linux-amd64
	rsync --progress $< aws-www:~/bin/www-kilabit

deploy-local:
	rsync ./www-kilabit $(GOBIN)/

serve:
	DEBUG=1 go run ./cmd/www-kilabit
