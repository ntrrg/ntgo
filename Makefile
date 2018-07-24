gofiles := $(shell find . -iname "*.go" -type f)
coverage_file := coverage.txt

.PHONY: all
all: build

.PHONY: benchmark
benchmark:
	go test -race -bench . -benchmem ./...

.PHONY: build
build:
	go build ./...

.PHONY: check
check: test lint coverage benchmark

.PHONY: ci
ci: test lint qa coverage benchmark

.PHONY: clean
clean:
	rm -f $(coverage_file)

.PHONY: coverage
coverage:
	@go test -race -coverprofile $(coverage_file) ./... > /dev/null
	go tool cover -func $(coverage_file)

.PHONY: coverage-web
coverage-web:
	@go test -race -coverprofile $(coverage_file) ./... > /dev/null
	go tool cover -html $(coverage_file)

.PHONY: deps
deps:
	@which gometalinter.v2 > /dev/null 2> /dev/null \
		|| (go get -u gopkg.in/alecthomas/gometalinter.v2 \
		&& gometalinter.v2 --install)

.PHONY: docs
docs:
	@echo "See http://localhost:6060/pkg/github.com/ntrrg/ntgo"
	godoc -http :6060 -play

.PHONY: format
format:
	gofmt -s -w -l $(gofiles)

.PHONY: install
install:
	go install -i ./...

.PHONY: lint
lint: deps
	gofmt -d -e -s $(gofiles)
	gometalinter.v2 --tests --fast ./...

.PHONY: lint-md
lint-md:
	@docker run --rm -itv "$$PWD":/files/ ntrrg/md-linter

.PHONY: qa
qa: deps
	CGO_ENABLED=0 gometalinter.v2 --tests --disable=interfacer ./...

.PHONY: test
test:
	go test -race -v ./...

