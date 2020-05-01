pkgName := ntgo
hugoPort := 1313
godocPort := 6060

goAllFiles := $(filter-out ./vendor/%, $(shell find . -iname "*.go" -type f))
goSrcFiles := $(shell go list -f "{{ \$$path := .Dir }}{{ range .GoFiles }}{{ \$$path }}/{{ . }} {{ end }}" ./...)
goTestFiles := $(shell go list -f "{{ \$$path := .Dir }}{{ range .TestGoFiles }}{{ \$$path }}/{{ . }} {{ end }} {{ range .XTestGoFiles }}{{ \$$path }}/{{ . }} {{ end }}" ./...)

.PHONY: all
all: build

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean:
	rm -rf dist/

.PHONY: doc
doc:
	@echo "Go to http://localhost:$(hugoPort)/en/projects/$(pkgName)/"
	@echo "Ir a http://localhost:$(hugoPort)/es/projects/$(pkgName)/"
	@docker run --rm -it \
		-e PORT=$(hugoPort) \
		-p $(hugoPort):$(hugoPort) \
		-v "$$PWD/.ntweb":/site/content/projects/$(pkgName)/ \
		ntrrg/ntweb:editing --port $(hugoPort)

.PHONE: doc-go
doc-go:
	@echo "Go to http://localhost:$(godocPort)/pkg/$(shell go list -m)/"
	godoc -http :$(godocPort) -play

# Development

coverage_file := coverage.txt
CI_TARGET ?= ./...

.PHONY: benchmark
benchmark:
	go test -v -bench . -benchmem $(CI_TARGET)

.PHONY: ca
ca:
	golangci-lint run

.PHONY: ci
ci: clean-dev test lint ca coverage benchmark build

.PHONY: ci-race
ci-race: clean-dev test-race lint ca coverage benchmark build

.PHONY: clean-dev
clean-dev: clean
	rm -rf $(coverage_file)

.PHONY: coverage
coverage: $(coverage_file)
	go tool cover -func $<

.PHONY: coverage-web
coverage-web: $(coverage_file)
	go tool cover -html $<

.PHONY: format
format:
	gofmt -s -w -l $(goAllFiles)

.PHONY: lint
lint:
	gofmt -d -e -s $(goAllFiles)

.PHONY: test
test:
	go test -v $(CI_TARGET)

.PHONY: test-race
test-race:
	go test -race -v $(CI_TARGET)

.PHONY: watch
watch:
	reflex -d "none" -r '\.go$$' -- $(MAKE) -s test lint

.PHONY: watch-race
watch-race:
	reflex -d "none" -r '\.go$$' -- $(MAKE) -s test-race lint

$(coverage_file): $(goSrcFiles) $(goTestFiles)
	go test -coverprofile $(coverage_file) $(CI_TARGET)

