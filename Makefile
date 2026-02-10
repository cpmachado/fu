MAIN=./cmd/fud

all: build
	@echo all built

build:
	 go build $(MAIN)

clean:
	@rm -rf fud target
	@echo all removed

lint:
	golangci-lint run ./...
	@echo all code is linted

format:
	gofmt -w -s .

format-check:
	gofmt -l .

run:
	go run $(MAIN)

test:
	@mkdir -p target
	go test -v -coverprofile=target/fud-coverage.out ./...

sbom: build
	@mkdir -p target/sbom
	cyclonedx-gomod bin -json -output ./target/sbom/fud.bom.json ./fud

.PHONY: all build clean lint format format-check run test sbom
