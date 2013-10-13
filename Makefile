BIN=$(GOPATH)/bin

$(BIN)/test-loglevel: *.go
	go test -c
	mv loglevel.test $(BIN)/test-loglevel

test: $(BIN)/test-loglevel
ifdef TEST
	$(BIN)/test-loglevel -test.run="$(TEST)"
else
	$(BIN)/test-loglevel -test.v
endif

format:
	gofmt -w *.go

lint:
	go get github.com/golang/lint/golint
	$(BIN)/golint *.go
	go vet

.PHONY: test format lint
