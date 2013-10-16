BIN=$(GOPATH)/bin

test: export GOPATH=
test:
	go test -v ./...

format:
	gofmt -w *.go

lint:
	go get github.com/golang/lint/golint
	$(BIN)/golint *.go
	go vet

.PHONY: test format lint
