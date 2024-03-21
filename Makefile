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
