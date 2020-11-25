.PHONY: all deploy deploy-local

all: serve

deploy:
	go generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o www-kilabit-linux-amd64 ./cmd/www-kilabit/
	rsync --progress www-kilabit-linux-amd64 personal-server:~/bin/www-kilabit

deploy-local:
	go generate
	go build ./cmd/www-kilabit
	rsync ./www-kilabit $(GOBIN)/

serve:
	ulimit -n 8192
	DEBUG=1 go run ./cmd/www-kilabit
