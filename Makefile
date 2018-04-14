VERSION = `git rev-parse HEAD`
DATE = `date --iso-8601=seconds`
LDFLAGS =  -X github.com/boreq/upspin-manifest/cmd.BuildCommit=$(VERSION)
LDFLAGS += -X github.com/boreq/upspin-manifest/cmd.BuildDate=$(DATE)

all: build

build:
	mkdir -p build
	go build -ldflags "$(LDFLAGS)" -o ./build/upspin-manifest ./cmd/upspin-manifest

doc:
	@echo "http://localhost:6060/pkg/github.com/boreq/upspin-manifest/"
	godoc -http=:6060

test:
	go test ./...

test-verbose:
	go test -v ./...

test-short:
	go test -short ./...

bench:
	go test -v -run=XXX -bench=. ./...

clean:
	rm -rf ./build

.PHONY: all build doc test test-verbose test-short bench clean
