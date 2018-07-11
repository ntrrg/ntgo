include config.mk

.PHONY: all
all: build

.PHONY: benchmark
benchmark:
	go test -bench . -benchmem ./...

.PHONY: build
build:
	go build -i ./...

.PHONY: build-docker
build-docker:
	docker build -t ntrrg/ntgo  .

.PHONY: ci
ci: test lint-go qa coverage benchmark

.PHONY: clean
clean:
	rm -f $(coverage_results)
	docker rm -f ntrrg/ntgo || true

.PHONY: coverage
coverage:
	@go test -covermode count -coverprofile $(coverage_results) ./... > /dev/null
	go tool cover -func $(coverage_results)

.PHONY: coverage-web
coverage-web:
	@go test -covermode count -coverprofile $(coverage_results) ./... > /dev/null
	go tool cover -html $(coverage_results)

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
lint: lint-md lint-go

.PHONY: lint-go
lint-go: deps
	gofmt -d -e -s $(gofiles)
	gometalinter.v2 --fast ./...

.PHONY: lint-md
lint-md:
	@docker run --rm -itv "$$PWD":/files/ $(mdlinter_image)

.PHONY: qa
qa: deps
	gometalinter.v2 ./...

.PHONY: test
test:
	go test -v ./...

