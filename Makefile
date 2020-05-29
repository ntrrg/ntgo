module := $(shell go list -m)
hugoPort := 1313
godocPort := 6060

goAllFiles := $(filter-out ./vendor/%, $(shell find . -iname "*.go" -type f))
goSrcFiles := $(shell go list -f "{{ \$$path := .Dir }}{{ range .GoFiles }}{{ \$$path }}/{{ . }} {{ end }}" ./...)
goTestFiles := $(shell go list -f "{{ \$$path := .Dir }}{{ range .TestGoFiles }}{{ \$$path }}/{{ . }} {{ end }}{{ range .XTestGoFiles }}{{ \$$path }}/{{ . }} {{ end }}" ./...)

.PHONY: all
all: build

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean: clean-dev

.PHONY: doc
doc:
	@echo "Go to http://localhost:$(hugoPort)/en/projects/$(basename $(module))/"
	@echo "Ir a http://localhost:$(hugoPort)/es/projects/$(basename $(module))/"
	@docker run --rm -it \
		-e PORT=$(hugoPort) \
		-p $(hugoPort):$(hugoPort) \
		-v "$$PWD/.ntweb":/site/content/projects/$(basename $(module))/ \
		ntrrg/ntweb:editing --port $(hugoPort)

.PHONE: doc-go
doc-go:
	@echo "Go to http://localhost:$(godocPort)/pkg/$(module)/"
	godoc -http :$(godocPort) -play

# Development

CI_TARGET ?= ./...
COVERAGE_FILE ?= coverage.txt

.PHONY: benchmark
benchmark:
	go test -v -bench . -benchmem -run none $(CI_TARGET)

.PHONY: ca
ca:
	golangci-lint run

.PHONY: ci
ci: test lint ca coverage benchmark build

.PHONY: ci-race
ci-race: test-race lint ca coverage benchmark build

.PHONY: clean-dev
clean-dev: clean
	rm -rf $(COVERAGE_FILE)

.PHONY: coverage
coverage: $(COVERAGE_FILE)
	go tool cover -func $<

.PHONY: coverage-web
coverage-web: $(COVERAGE_FILE)
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

$(COVERAGE_FILE): $(goSrcFiles) $(goTestFiles)
	go test -coverprofile $@ $(CI_TARGET)

