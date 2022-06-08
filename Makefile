.PHONY: all serve

all: serve

serve:
	go run ./cmd/www-kilabit -env=dev

## Local tasks.

.PHONY: local-setup local-deploy

local-deploy:
	go generate
	go build ./cmd/www-kilabit
	rsync --progress ./www-kilabit dev.local:/data/bin/

##---- Remote tasks.

.PHONY: deploy

deploy:
	go generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o www-kilabit-linux-amd64 ./cmd/www-kilabit/
	rsync --progress www-kilabit-linux-amd64 www-kilabit:/data/app/bin/www-kilabit
