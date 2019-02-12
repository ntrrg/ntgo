// +build mage

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

var (
	goFiles    = getGoFiles()
	goSrcFiles = getGoSrcFiles()
)

var Default = Build

func Build() error {
	return sh.RunV("go", "build", "./...")
}

// Development

var (
	coverageFile = "coverage.txt"
)

func Benchmark() error {
	return sh.RunV("go", "test", "-bench", ".", "-benchmem", "./...")
}

func CI() {
	mg.SerialDeps(Lint, QA, Test, Coverage.Default, Benchmark, Build)
}

func Clean() {
	sh.Rm(coverageFile)
}

type Coverage mg.Namespace

func (Coverage) Default() error {
	mg.Deps(CoverageFile)
	return sh.RunV("go", "tool", "cover", "-func", coverageFile)
}

func (Coverage) Web() error {
	mg.Deps(CoverageFile)
	return sh.RunV("go", "tool", "cover", "-html", coverageFile)
}

func CoverageFile() error {
	if run, err := target.Path(coverageFile, goFiles...); !run || err != nil {
		return err
	}

	return Test()
}

func Docs() {
	sh.RunV("godoc", "-http", ":6060", "-play")
}

func Format() error {
	args := []string{"-s", "-w", "-l"}
	args = append(args, goFiles...)
	return sh.RunV("gofmt", args...)
}

func Lint() error {
	args := []string{"-d", "-e", "-s"}
	args = append(args, goFiles...)
	return sh.RunV("gofmt", args...)
}

func QA() error {
	return sh.RunV("golangci-lint", "run")
}

func Test() error {
	return sh.RunV("go", "test", "-race", "-coverprofile", coverageFile, "-v", "./...")
}

// Helpers

func getGoFiles() []string {
	var goFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "vendor/") {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})

	return goFiles
}

func getGoSrcFiles() []string {
	var goSrcFiles []string

	for _, path := range goFiles {
		if !strings.HasSuffix(path, "_test.go") {
			continue
		}

		goSrcFiles = append(goSrcFiles, path)
	}

	return goSrcFiles
}
