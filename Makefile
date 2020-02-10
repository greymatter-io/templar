VERSIONS_PACKAGE := github.com/greymatter-io/templar/versions

COMMIT := $(shell git rev-parse --verify --short HEAD 2> /dev/null || echo "UNKNOWN")
COMMIT_FLAG := -X $(VERSIONS_PACKAGE).commit=$(COMMIT)

VERSION := $(shell cat VERSION || echo "UNKNOWN")
VERSION_FLAG := -X $(VERSIONS_PACKAGE).version=$(VERSION)

.PHONY: build
build: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 go build -o bin/templar -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: build.docker
build.docker: build.linux
	@echo "--> Building docker..."
	@docker build -f Dockerfile -t "greymatterio/templar:$(VERSION)" .

.PHONY: build.linux
build.linux: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/templar.linux -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: build.macos
build.macos: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/templar.macos -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: test
test: vendor
	@echo "--> Running tests..."
	@CGO_ENABLED=0 go test -v --coverprofile=./coverage/c.out --mod=vendor ./...

.PHONY: vendor
vendor:
	@echo "--> Vendoring dependencies..."
	@CGO_ENABLED=0 go mod vendor
