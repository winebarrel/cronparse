.PHONY: all
all: vet lint test

.PHONY: vet
vet:
	go vet -composites=false -structtag=false ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...
