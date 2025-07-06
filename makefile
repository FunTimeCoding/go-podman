.DEFAULT_GOAL := all

# Source: https://github.com/containers/podman/blob/v5.5.0/Makefile#L60
REMOTE_TAGS ?= remote exclude_graphdriver_btrfs btrfs_noversion containers_image_openpgp

all: lint test

tool:
	@go install gotest.tools/gotestsum@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/gobuild@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/golint@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/goupdate@latest

lint:
	@golint --fix
	@golangci-lint run --build-tags "${REMOTE_TAGS}"

test:
	@gotestsum --format standard-quiet -- -tags "${REMOTE_TAGS}" ./...

update:
	@goupdate

update-library:
	@goupdate --exclusive funtimecoding

install:
	@gobuild --all --copy-to-bin --build-tags "${REMOTE_TAGS}"

postgres:
	@go run -tags "${REMOTE_TAGS}" cmd/example/postgres/main.go

build-github:
	@go build -v -tags "${REMOTE_TAGS}" ./...
