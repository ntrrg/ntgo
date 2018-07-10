coverage_results := coverage.out
mdlinter_image := ntrrg/md-linter

gofiles := $(shell find . -iname "*.go" -type f)
srcfiles := $(shell echo $(gofiles) | sed -re "s/(\/?\w)+_test.go//g")

