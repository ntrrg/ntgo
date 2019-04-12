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

var Default = Build

func Build() error {
	return sh.RunV("go", "build", "./...")
}

// Development

var (
	goFiles      = getGoFiles()
	coverageFile = "coverage.txt"
)

func Benchmark() error {
	return sh.RunV("go", "test", "-bench", ".", "-benchmem", "./...")
}

func CI() {
	mg.SerialDeps(Lint, QA, Test.Race, Coverage.Default, Benchmark, Build)
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

	return Test{}.Default()
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

type Test mg.Namespace

func (Test) Default() error {
	return sh.RunV("go", "test", "-coverprofile", coverageFile, "-v", "./...")
}

func (Test) Race() error {
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
