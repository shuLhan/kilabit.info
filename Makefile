.PHONY: all embed serve find-broken-symlinks

all: serve

embed:
	go run ./cmd/www-kilabit embed

serve:
	go run ./cmd/www-kilabit -dev

find-broken-symlinks:
	@echo ">>> Finding broken symlinks ..."
	@broken=$$(find . -xtype l); \
	echo $$broken; \
	if [[ "$$broken" != "" ]]; then exit 1; fi

##---- Local tasks.

.PHONY: local-setup local-deploy

local-deploy: find-broken-symlinks embed
	go build ./cmd/www-kilabit
	rsync --progress ./www-kilabit dev.local:/data/bin/

##---- Remote tasks.

.PHONY: deploy

deploy: find-broken-symlinks embed
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o www-kilabit-linux-amd64 ./cmd/www-kilabit/
	rsync --progress www-kilabit-linux-amd64 kilabit.info:/data/app/bin/www-kilabit
