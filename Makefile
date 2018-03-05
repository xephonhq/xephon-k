VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/xk
	go install -ldflags "$(FLAGS)" ./cmd/xkctl

.PHONY: generate
generate:
	gommon generate -v

.PHONY: fmt
fmt:
	gofmt -d -l -w ./cmd/xk ./cmd/xkctl ./xk

.PHONY: test
test:
	go test -v -cover ./xk/...

.PHONY: loc
loc:
	cloc --exclude-dir=vendor,.idea,playground --exclude-list-file=script/cloc_exclude.txt .