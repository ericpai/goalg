.PHONY: fmt lint linter-install test

fmt:
	goimports -local -l -w app cmd tools

lint:
	golangci-lint run ./

linter-install:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(GOPATH)/bin v1.20.0

test:
	go test -v ./... -coverprofile=coverage.txt -covermode=atomic

