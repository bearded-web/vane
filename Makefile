.PHONY: all test lint vet fmt travis


export FIXTURESPATH=$(shell pwd)/fixtures

all: test lint vet fmt


travis: all


test:
	@echo "+ $@"
	@go test -cover ./...


lint:
	@echo "+ $@"
	@test -z "$$(golint ./... | grep -v Godeps/_workspace/src/ | tee /dev/stderr)"


vet:
	@echo "+ $@"
	@go vet ./...


fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . | grep -v Godeps/_workspace/src/ | tee /dev/stderr)" || \
		echo "+ please format Go code with 'gofmt -s'"
