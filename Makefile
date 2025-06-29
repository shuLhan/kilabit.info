.PHONY: build
build: embed
	go build -o _bin/ ./cmd/www-kilabit/

.PHONY: embed
embed:
	go run ./cmd/www-kilabit embed

.PHONY: serve
serve:
	go run ./cmd/www-kilabit -dev -address=127.0.0.1:17000

##---- Install to local GOBIN

.PHONY: install
install: embed
	go install ./cmd/...

##---- Remote tasks.

.PHONY: deploy
deploy: CGO_ENABLED=0
deploy: GOOS=linux
deploy: GOARCH=amd64
deploy: build
	rsync --progress _bin/www-kilabit kilabit.info:/data/app/bin/www-kilabit

##---- Task on webhook events.

.PHONY: on-webhook
on-webhook: CGO_ENABLED=0
on-webhook: GOOS=linux
on-webhook: GOARCH=amd64
on-webhook: build
	sudo rsync --progress _bin/www-kilabit /data/app/bin/www-kilabit

##---- Scan broken links using jarink.\
## -ignore-status
##  403 - Forbidden, usually pages that require login.
##  418 - Teapot, usually pages blocked from scanned by AI bot, like
##        sr.ht website.
##  429 - Too many requests.

.PHONY: jarink.brokenlinks
jarink.brokenlinks:
	jarink \
		-ignore-status=403,418,429 \
		-insecure \
		-past-result=jarink_brokenlinks.json \
		-verbose \
		brokenlinks \
		https://kilabit.home.local \
		> jarink_brokenlinks.json
