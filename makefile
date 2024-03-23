.DEFAULT_GOAL := all

# Source: https://github.com/containers/podman/blob/v4.8.3/Makefile#L51
REMOTE_TAGS ?= remote exclude_graphdriver_btrfs btrfs_noversion exclude_graphdriver_devicemapper containers_image_openpgp

update:
	go install gotest.tools/gotestsum@latest
	for ELEMENT in $$(go list -f "{{if not (or .Main .Indirect)}}{{.Path}}{{end}}" -m all); do echo $${ELEMENT}; go get $${ELEMENT}; done
	go get github.com/docker/docker@v24.0.9
	go get github.com/containers/common@v0.56.0
	#go get github.com/containers/podman/v4@v4.7.2
	go mod tidy

update-library:
	GOPROXY=direct go get github.com/funtimecoding/go-library
	go mod tidy

lint:
	golangci-lint run --build-tags "${REMOTE_TAGS}"

test:
	gotestsum --format standard-quiet -- -tags "${REMOTE_TAGS}" ./...
