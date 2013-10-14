BIN=$(GOPATH)/bin

test:
	go test

format:
	gofmt -w *.go

lint:
	go get github.com/golang/lint/golint
	$(BIN)/golint *.go
	go vet

.PHONY: test format lint
