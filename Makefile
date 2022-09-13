GO ?= go

module := $(shell $(GO) list -m)
PROJECT ?= $(notdir $(module))

GODOC_PORT ?= 6060
NTWEB_PORT ?= 1313

goFiles := $(shell find . -iname "*.go" -type f | grep -v "/_" | grep -v "^\./vendor")
goFilesSrc := $(shell $(GO) list -f '{{ range .GoFiles }}{{ $$.Dir }}/{{ . }} {{ end }}' ./...)
goFilesTest := $(shell $(GO) list -f "{{ range .TestGoFiles }}{{ $$.Dir }}/{{ . }} {{ end }}{{ range .XTestGoFiles }}{{ $$.Dir }}/{{ . }} {{ end }}" ./...)

.PHONY: all
all: build

.PHONY: build
build:
	$(GO) build "./..."

.PHONY: clean
clean:

# Development

BENCHMARK_COUNT ?= 1
BENCHMARK_FILE ?= benchmarks-dev.txt
BENCHMARK_WEB_FILE := $(shell mktemp -u)-$(PROJECT).html
COVERAGE_FILE ?= coverage-dev.txt
CPUPROFILE ?= cpu.prof
MEMPROFILE ?= mem.prof
TARGET_FUNC ?= .
TARGET_PKG ?= ./...
WATCH_TARGET ?= test

ifneq "$(notdir $(TARGET_PKG))" "..."
	profileFlags := -cpuprofile "$(CPUPROFILE)" -memprofile "$(MEMPROFILE)"
endif

.PHONY: benchmark
benchmark:
	$(GO) test -v -run none \
		-bench "$(TARGET_FUNC)" -benchmem -count $(BENCHMARK_COUNT) \
		$(profileFlags) \
		"$(TARGET_PKG)" | tee "$(BENCHMARK_FILE)"

.PHONY: benchmark-check
benchmark-check: benchmarks.txt $(BENCHMARK_FILE)
	benchstat "$<" "$(BENCHMARK_FILE)"

.PHONY: benchmark-web
benchmark-web: benchmarks.txt $(BENCHMARK_FILE)
	benchstat -html "$<" "$(BENCHMARK_FILE)" > "$(BENCHMARK_WEB_FILE)"
	exo-open "$(BENCHMARK_WEB_FILE)"

define benchmarks_file
	BENCHMARK_FILE="$(1)" $(MAKE) -s benchmark
endef

benchmarks.txt:
	$(call benchmarks_file,$@)

benchmarks-%.txt:
	$(call benchmarks_file,$@)

.PHONY: ca
ca:
	golangci-lint run

.PHONY: ca-fast
ca-fast:
	golangci-lint run --fast

.PHONY: ci
ci: build test lint ca

.PHONY: ci-race
ci-race: build test-race lint ca

.PHONY: clean-dev
clean-dev: clean
	rm -rf "$(CPUPROFILE)" "$(MEMPROFILE)" benchmarks-*.txt coverage-*.txt *.test

.PHONY: coverage
coverage: $(COVERAGE_FILE)
	$(GO) tool cover -func "$(COVERAGE_FILE)"

.PHONY: coverage-check
coverage-check: coverage.txt $(COVERAGE_FILE)
	#coverstat "$<" "$(COVERAGE_FILE)"

.PHONY: coverage-web
coverage-web: $(COVERAGE_FILE)
	$(GO) tool cover -html "$(COVERAGE_FILE)"

define coverage_file
	COVERAGE_FILE="$(1)" $(MAKE) -s test
endef

coverage.txt:
	$(call coverage_file,$@)

coverage-%.txt:
	$(call coverage_file,$@)

.PHONY: doc
doc:
	@echo "Go to http://localhost:$(NTWEB_PORT)/en/projects/$(PROJECT)/"
	@docker run --rm -it \
		-e "PORT=$(NTWEB_PORT)" \
		-p "$(NTWEB_PORT):$(NTWEB_PORT)" \
		-v "$$PWD/.ntweb:/site/content/projects/$(PROJECT)/" \
		ntrrg/ntweb:editing --port "$(NTWEB_PORT)"

.PHONY: format
format:
	gofmt -s -w -l $(goFiles)

.PHONY: fuzz
fuzz:
	$(GO) test -v -run none -fuzz "$(TARGET_FUNC)" "$(TARGET_PKG)"

.PHONY: godoc
godoc:
	@echo "Go to http://localhost:$(GODOC_PORT)/pkg/$(module)/"
	godoc -http ":$(GODOC_PORT)" -play
	#GOPROXY=$(shell go env GOPROXY) pkgsite -http ":$(GODOC_PORT)" -cache -proxy

.PHONY: lint
lint:
	gofmt -d -e -s $(goFiles)

.PHONY: test
test:
	$(GO) test -v \
		-run "$(TARGET_FUNC)" \
		-coverprofile "$(COVERAGE_FILE)" \
		$(profileFlags) \
		"$(TARGET_PKG)"

.PHONY: test-race
test-race:
	$(GO) test -v -race \
		-run "$(TARGET_FUNC)" \
		-coverprofile "$(COVERAGE_FILE)" \
		$(profileFlags) \
		"$(TARGET_PKG)"

.PHONY: watch
watch:
	reflex -d "none" -r '\.go$$' -s -- $(MAKE) -s $(WATCH_TARGET)
