.DEFAULT_GOAL := all

# Source: https://github.com/containers/podman/blob/v5.1.0-rc1/Makefile#L51
REMOTE_TAGS ?= remote exclude_graphdriver_btrfs btrfs_noversion exclude_graphdriver_devicemapper containers_image_openpgp

all: lint test

tool:
	@go install gotest.tools/gotestsum@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/gobuild@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/golint@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/goupdate@latest

lint:
	@golangci-lint run --build-tags "${REMOTE_TAGS}"

test:
	@gotestsum --format standard-quiet -- -tags "${REMOTE_TAGS}" ./...

update:
	@goupdate --downgrade github.com/docker/docker@v26.1.5 \
    --downgrade github.com/containers/common@v0.58.3

update-library:
	@goupdate --exclusive funtimecoding

postgres:
	@go run -tags "${REMOTE_TAGS}" cmd/example/postgres/main.go
