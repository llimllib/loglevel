BIN=$(GOPATH)/bin

$(BIN)/test-log: *.go
	go test -c
	mv log.test $(BIN)/test-log

test: $(BIN)/test-log
ifdef TEST
	$(BIN)/test-log -test.run="$(TEST)"
else
	$(BIN)/test-log -test.v
endif

format:
	gofmt -w *.go

lint:
	go get github.com/golang/lint/golint
	$(BIN)/golint *.go
	go vet

.PHONY: test format lint
