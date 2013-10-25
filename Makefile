BIN=$(GOPATH)/bin

# The testfatal/* tests need to be in package main so that I can "go run" them.
# They also need to import loglevel. If the source tree is contained within
# a GOPATH, I must use "import loglevel" instead of "import ../../loglevel".
# Most of the time I'm working in the repo which is not in a GOPATH, so instead
# of leaving it as "import gopath", I'll just make sure that we're not inside
# a GOPATH when the tests are run by blanking out the GOPATH variable. This
# makes Travis happy and doesn't annoy me during dev.
#
# TODO: figure out how to kill this hack
test: GOPATH=
test:
	go test -v ./...

format:
	gofmt -w *.go

lint:
	go get github.com/golang/lint/golint
	$(BIN)/golint *.go
	go vet

.PHONY: test format lint
