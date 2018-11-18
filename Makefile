gofiles := $(shell find . -iname "*.go" -type f)
make_bin := /tmp/$(shell basename "$$PWD")-bin

.PHONY: all
all: build

.PHONY: build
build:
	go build ./...

.PHONY: docs
docs:
	@echo "See http://localhost:6060/pkg/github.com/ntrrg/ntgo"
	godoc -http :6060 -play

.PHONY: install
install:
	go install -i ./...

# Development

coverage_file := coverage.txt

.PHONY: benchmark
benchmark:
	go test -race -bench . -benchmem ./...

.PHONY: ci
ci: test lint qa coverage benchmark

.PHONY: clean
clean-dev:
	rm -f $(coverage_file)

.PHONY: coverage
coverage: test
	go tool cover -func $(coverage_file)

.PHONY: coverage-web
coverage-web: test
	go tool cover -html $(coverage_file)

.PHONY: deps-dev
deps-dev: $(make_bin)/gometalinter
	go get -u -v golang.org/x/lint/golint

.PHONY: format
format:
	gofmt -s -w -l $(gofiles)

.PHONY: lint
lint:
	gofmt -d -e -s $(gofiles)
	golint ./...

.PHONY: lint-md
lint-md:
	@docker run --rm -itv "$$PWD":/files/ ntrrg/md-linter

.PHONY: qa
qa: $(make_bin)/gometalinter
	PATH="$(make_bin):$$PATH" CGO_ENABLED=0 gometalinter --tests ./ ./api/... ./pkg/...

.PHONY: test
test:
	go test -race -coverprofile $(coverage_file) -v ./...

$(make_bin)/gometalinter:
	mkdir -p $(make_bin)
	wget -cO /tmp/gometalinter.tar.gz 'https://storage.nt.web.ve/_/software/linux/gometalinter-2.0.11-linux-amd64.tar.gz' || wget -cO /tmp/gometalinter.tar.gz 'https://github.com/alecthomas/gometalinter/releases/download/v2.0.11/gometalinter-2.0.11-linux-amd64.tar.gz'
	tar -xf /tmp/gometalinter.tar.gz -C /tmp/
	cp -a $$(find /tmp/gometalinter-2.0.11-linux-amd64/ -type f -executable) $(make_bin)/

