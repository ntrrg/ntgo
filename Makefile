include config.mk

.PHONY: all
all:

.PHONY: benchmark
benchmark:
	@go test -bench . -benchmem ./...

.PHONY: ci
ci: test qa lint-go coverage benchmark

.PHONY: clean
clean:
	rm -f $(coverage_results)

.PHONY: coverage
coverage: $(coverage_results)
	@go tool cover -func $<

.PHONY: coverage-web
coverage-web: $(coverage_results)
	@go tool cover -html $<

.PHONY: dev-deps
dev-deps:
	@which gometalinter.v2 > /dev/null 2> /dev/null \
		|| (go get -u gopkg.in/alecthomas/gometalinter.v2 \
		&& gometalinter.v2 --install)

.PHONY: godoc
godoc:
	godoc -http :6060 -play

.PHONY: format
format:
	@gofmt -s -w -l $(gofiles)

.PHONY: lint
lint: lint-md lint-go

.PHONY: lint-go
lint-go: dev-deps
	gofmt -d -e -s $(gofiles)
	golint ./...

.PHONY: lint-md
lint-md:
	@docker run --rm -itv "$$PWD":/files/ $(mdlinter_image)

.PHONY: qa
qa: dev-deps
	@gometalinter.v2 ./...

.PHONY: test
test:
	@go test -v ./...

$(coverage_results): $(go_files)
	@mkdir -p $$(dirname $@)
	@go test -covermode count -coverprofile $@ ./... > /dev/null

