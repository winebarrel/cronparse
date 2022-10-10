.PHONY: all
all: vet test build

.PHONY: build
build:
	go build ./cmd/cronnext

.PHONY: vet
vet:
	go vet -composites=false -structtag=false ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f cronnext cronnext.exe
