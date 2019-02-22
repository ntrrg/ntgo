// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package os_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	ntos "github.com/ntrrg/ntgo/os"
)

func TestCp_noSource(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-ns")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	dst := filepath.Join(dir, "dest.txt")
	if err := ntos.Cp(dst); err == nil {
		t.Error("ntos.Cp succeeded without sources")
	}
}

func TestCp_nonExistentSource(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-nes")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "non-existent.txt")
	dst := filepath.Join(dir, "dest.txt")
	if err := ntos.Cp(dst, src); err == nil {
		t.Error("ntos.Cp succeeded copying non existent source")
	}
}

func TestCp_multipleNonExistentElementsInDestPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-mneeidp")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "file.txt")
	dst := filepath.Join(dir, "non/existent/path")
	if err := ntos.Cp(dst, src); err == nil {
		msg := "ntos.Cp succeeded copying to a multiple non existent elements path"
		t.Errorf(msg)
	}
}

func TestCp_badPermissions(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-bp")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "file.txt")
	dst := filepath.Join(dir, "forbidden")

	if err := os.Mkdir(dst, 0000); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err == nil {
		t.Error("ntos.Cp succeeded copying to a forbidden path")
	}

	src = filepath.Join(dir, "directory")
	dst = filepath.Join(dir, "other-directory")

	if err := os.Mkdir(src, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(src, "forbidden"), 0000); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(dst, 0755); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err == nil {
		t.Error("ntos.Cp succeeded copying elements with bad permissions")
	}
}

func TestCp_fileToNonExistentFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-f2nef")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "file.txt")
	dst := filepath.Join(dir, "non-existent.txt")

	if err := ioutil.WriteFile(src, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err != nil {
		msg := "ntos.Cp can't copy %s to non existent %s: %v"
		t.Errorf(msg, src, dst, err)
	}

	if err := compareFiles(dst, src); err != nil {
		t.Errorf("invalid copy: %v", err)
	}
}

func TestCp_fileToFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-f2f")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "file.txt")
	dst := filepath.Join(dir, "other-file.txt")

	if err := ioutil.WriteFile(src, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile(dst, []byte("other file"), 0644); err != nil {
		t.Fatalf("can't create file %s: %v", dst, err)
	}

	if err := ntos.Cp(dst, src); err != nil {
		t.Errorf("ntos.Cp can't copy %s to %s: %v", src, dst, err)
	}

	if err := compareFiles(dst, src); err != nil {
		t.Errorf("invalid copy: %v", err)
	}
}

func TestCp_fileToDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-f2d")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "file.txt")
	dst := filepath.Join(dir, "directory")

	if err := ioutil.WriteFile(src, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(dst, 0744); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err != nil {
		t.Errorf("ntos.Cp can't copy %s to %s/: %v", src, dst, err)
	}

	dst = filepath.Join(dst, filepath.Base(src))
	if err := compareFiles(dst, src); err != nil {
		t.Errorf("invalid copy: %v", err)
	}
}

func TestCp_directoryToFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-d2f")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "directory")
	dst := filepath.Join(dir, "file.txt")

	if err := os.Mkdir(src, 0744); err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile(dst, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err == nil {
		t.Error("ntos.Cp succeeded copying directory to file")
	}
}

func TestCp_directoryToNonExistentDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-d2ned")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "directory")
	dst := filepath.Join(dir, "non-existent")

	if err := os.Mkdir(src, 0744); err != nil {
		t.Fatal(err)
	}

	file := filepath.Join(src, "file.txt")
	if err := ioutil.WriteFile(file, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err != nil {
		t.Errorf("ntos.Cp can't copy %s/ to %s/: %v", src, dst, err)
	}

	if err := compareDirs(dst, src); err != nil {
		t.Errorf("invalid copy: %v", err)
	}
}

func TestCp_directoryToDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-d2d")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "directory")
	dst := filepath.Join(dir, "other-directory")

	if err := os.Mkdir(src, 0744); err != nil {
		t.Fatal(err)
	}

	file := filepath.Join(src, "file.txt")
	if err := ioutil.WriteFile(file, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(src, "subdirectory"), 0744); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(dst, 0744); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err != nil {
		t.Errorf("ntos.Cp can't copy %s/ to %s/: %v", src, dst, err)
	}

	dst = filepath.Join(dst, filepath.Base(src))
	if err := compareDirs(dst, src); err != nil {
		t.Errorf("invalid copy: %v", err)
	}
}

func TestCp_directoryToItSelf(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-d2is")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "directory")
	dst := filepath.Join(dir, "directory/inside")

	if err := os.Mkdir(src, 0744); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src); err == nil {
		t.Error("ntos.Cp succeeded copying directory to itself")
	}
}

func TestCp_multipleElementsToFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-m2f")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	dst := filepath.Join(dir, "other-file.txt")

	if err := ioutil.WriteFile(dst, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, "other-file.txt", "directory"); err == nil {
		t.Error("ntos.Cp succeeded copying multiple elements to a file")
	}
}

func TestCp_multipleElementsToNonExistentDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-me2ned")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	dst := filepath.Join(dir, "non/existent/directory")
	if err := ntos.Cp(dst, "file.txt", "directory"); err == nil {
		msg := "ntos.Cp succeeded copying multiple elements to a non existent dir"
		t.Errorf(msg)
	}
}

func TestCp_multipleElementsToDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "ntgo-os-cp-me2d")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "file.txt")
	src2 := filepath.Join(dir, "directory")
	dst := filepath.Join(dir, "other-directory")

	if err := ioutil.WriteFile(src, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(src2, 0744); err != nil {
		t.Fatal(err)
	}

	file := filepath.Join(src2, "file.txt")
	if err := ioutil.WriteFile(file, []byte("hello, world!"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(dst, 0744); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Cp(dst, src, src2); err != nil {
		t.Errorf("ntos.Cp can't copy %s and %s/ to %s/: %v", src, src2, dst, err)
	}

	dst2 := filepath.Join(dst, filepath.Base(src))
	if err := compareFiles(dst2, src); err != nil {
		t.Errorf("invalid copy: %v", err)
	}

	dst2 = filepath.Join(dst, filepath.Base(src2))
	if err := compareDirs(dst2, src2); err != nil {
		t.Errorf("invalid copy: %v", err)
	}
}

func compareDirs(dst, src string) error {
	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		dstfile := filepath.Clean(strings.Replace(path, src, dst, 1))
		return compareFiles(dstfile, path)
	}

	return filepath.Walk(src, fn)
}

func compareFiles(dst, src string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}

	dfi, err := os.Stat(dst)
	if err != nil {
		return err
	}

	if sfi.Size() != dfi.Size() {
		msg := "files have different sizes: %d (%s) and %d (%s)"
		return fmt.Errorf(msg, sfi.Size(), src, dfi.Size(), dst)
	}

	if sfi.Mode() != dfi.Mode() {
		msg := "files have different permissions: %#4o (%s) and %#4o (%s)"
		return fmt.Errorf(msg, sfi.Mode(), src, dfi.Mode(), dst)
	}

	sbytes, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("can't read %s: %v", src, err)
	}

	dbytes, err := ioutil.ReadFile(dst)
	if err != nil {
		return fmt.Errorf("can't read %s: %v", dst, err)
	}

	if !bytes.Equal(sbytes, dbytes) {
		return errors.New("files have different contents")
	}

	return nil
}

func tmpDir(name string) (string, error) {
	return ioutil.TempDir("", "ntgo-os-"+name)
}
