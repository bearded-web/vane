.PHONY: all test lint vet fmt travis


export FIXTURESPATH=$(shell pwd)/fixtures

all: test vet fmt


travis: deps all


deps:
	@go get -t ./...


test:
	@echo "+ $@"
	@go test -cover -v ./...


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
