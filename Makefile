.PHONY: all
all: serve

.PHONY: find-broken-symlinks
find-broken-symlinks:
	@echo ">>> Finding broken symlinks ..."
	@broken=$$(find . -xtype l); \
	echo $$broken; \
	if [[ "$$broken" != "" ]]; then exit 1; fi



.PHONY: embed
embed: find-broken-symlinks
	go run ./cmd/www-kilabit embed

.PHONY: build
build: embed
	go build ./cmd/www-kilabit/

.PHONY: serve
serve:
	go run ./cmd/www-kilabit -dev

##---- Local tasks.

.PHONY: local-deploy
local-deploy: build
	rsync --progress ./www-kilabit dev.local:/data/bin/

##---- Remote tasks.

.PHONY: deploy
deploy: CGO_ENABLED=0
deploy: GOOS=linux
deploy: GOARCH=amd64
deploy: build
	rsync --progress www-kilabit kilabit.info:/data/app/bin/www-kilabit
