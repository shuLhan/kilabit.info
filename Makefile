.PHONY: all serve find-broken-symlinks

all: serve

serve:
	go run ./cmd/www-kilabit -env=dev

find-broken-symlinks:
	@echo ">>> Finding broken symlinks ..."
	@broken=$$(find . -xtype l); \
	echo $$broken; \
	if [[ "$$broken" != "" ]]; then exit 1; fi

##---- Local tasks.

.PHONY: local-setup local-deploy

local-deploy: find-broken-symlinks
	go generate
	go build ./cmd/www-kilabit
	rsync --progress ./www-kilabit dev.local:/data/bin/

##---- Remote tasks.

.PHONY: deploy

deploy: find-broken-symlinks
	go generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o www-kilabit-linux-amd64 ./cmd/www-kilabit/
	rsync --progress www-kilabit-linux-amd64 www-kilabit:/data/app/bin/www-kilabit
