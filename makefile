.DEFAULT_GOAL := all

# Source: https://github.com/containers/podman/blob/v4.5.1/Makefile#L50
REMOTE_TAGS ?= remote exclude_graphdriver_btrfs btrfs_noversion exclude_graphdriver_devicemapper containers_image_openpgp

update:
	go install gotest.tools/gotestsum@latest
	GOPROXY=direct go get $$(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go mod tidy

update-library:
	GOPROXY=direct go get github.com/funtimecoding/go-library
	go mod tidy

lint:
	golangci-lint run --build-tags "${REMOTE_TAGS}"

test:
	gotestsum --format standard-quiet -- -tags "${REMOTE_TAGS}" ./...
