
export GO111MODULE=on

.PHONY: test
test:
	go test ./pkg/... ./cmd/... ./options/... -coverprofile cover.out

.PHONY: bin
bin: fmt vet
	go build -o bin/etcdviewer github.com/double12gzh/etcdviewer/cmd/etcdviewer

.PHONY: fmt
fmt:
	go fmt ./pkg/... ./cmd/... ./options/...

.PHONY: vet
vet:
	go vet ./pkg/... ./cmd/... ./options/...