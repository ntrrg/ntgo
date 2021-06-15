// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package os_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	ntos "go.ntrrg.dev/ntgo/os"
)

func TestCopy(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "source.txt")
	dst := filepath.Join(dir, "destination.txt")

	data := []byte("hello, world!")
	if err := os.WriteFile(src, data, 0o600); err != nil {
		t.Fatal(err)
	}

	if err := ntos.Copy(dst, src); err != nil {
		t.Fatalf("Copy failed to copy a valid file: %v", err)
	}

	src, dst = dir, dir+"2"
	if err := ntos.Copy(dst, src); err != nil {
		t.Fatalf("Copy failed to copy a valid directory: %v", err)
	}

	defer os.RemoveAll(dst)
}

func TestCopy_empty(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy_empty")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	dst := filepath.Join(dir, "destination.txt")
	if err := ntos.Copy(dst, ""); err == nil {
		t.Error("Copy succeeded without source")
	}
}

func TestCopyDir(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy-dir")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "source")
	dst := filepath.Join(dir, "destination")

	if err := os.Mkdir(src, 0o700); err != nil {
		t.Fatal(err)
	}

	data := []byte("hello, world!")
	file := filepath.Join(src, "file.txt")

	if err := os.WriteFile(file, data, 0o600); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(src, "subdirectory"), 0o700); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(dst, 0o700); err != nil {
		t.Fatal(err)
	}

	if err := ntos.CopyDir(dst, src, 0o700); err != nil {
		t.Fatalf("CopyDir failed to copy a valid directory: %v", err)
	}

	if err := compareDirs(dst, src); err != nil {
		t.Fatal(err)
	}
}

func TestCopyDir_empty(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy-dir_empty")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	dst := filepath.Join(dir, "dest")
	if err := ntos.CopyDir(dst, "", 0o700); err == nil {
		t.Error("CopyDir succeeded without source")
	}
}

func TestCopyDir_insideItself(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy-dir_inside")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "source")
	dst := filepath.Join(src, "subdir")

	if err := os.Mkdir(src, 0o700); err != nil {
		t.Fatal(err)
	}

	if err := ntos.CopyDir(dst, src, 0o700); err == nil {
		msg := "CopyDir succeeded copying a directory into itself"

		t.Errorf(msg)
	}
}

func TestCopyDir_multipleNonExistentElementsInDestPath(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy-dir_mneeidp")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "source")
	dst := filepath.Join(dir, "non/existent/path")

	if err := os.Mkdir(src, 0o700); err != nil {
		t.Fatal(err)
	}

	if err := ntos.CopyDir(dst, src, 0o700); err == nil {
		msg := "CopyDir succeeded copying to a multiple non existent elements path"
		t.Errorf(msg)
	}
}

func TestCopyError_Error(t *testing.T) {
	t.Parallel()

	src := "source.txt"
	dst := "destination.txt"
	reason := "can't find the source file"
	err := ntos.NewCopyError(src, dst, reason, os.ErrNotExist)
	got := err.Error()
	want := "cannot copy " + src + " to " + dst + ", " + reason
	want += ": " + os.ErrNotExist.Error()

	if got != want {
		t.Errorf("invalid error. got: %s, want: %s", got, want)
	}
}

func TestCopyError_Unwrap(t *testing.T) {
	t.Parallel()

	src := "source.txt"
	dst := "destination.txt"
	reason := "can't find the source file"
	err := ntos.NewCopyError(src, dst, reason, os.ErrNotExist)

	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("invalid wrapped error. got: %v, want: %v", err, os.ErrNotExist)
	}

	reason = "cannot copy a directory into itself"
	err = ntos.NewCopyError(src, dst, reason, nil)

	if errors.Unwrap(err) != nil {
		t.Error("invalid error, should not be wrapping another error")
	}
}

func TestCopyFile(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy-file")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "source.txt")
	dst := filepath.Join(dir, "destination.txt")

	if err := os.WriteFile(src, []byte("hello, world!"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := ntos.CopyFile(dst, src, 0o600); err != nil {
		t.Fatalf("CopyFile failed to copy a valid file: %v", err)
	}

	if err := compareFiles(dst, src); err != nil {
		t.Fatal(err)
	}
}

func TestCopyFile_directory2file(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy-file_directory2file")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "source")
	dst := filepath.Join(dir, "destination.txt")

	if err := os.Mkdir(src, 0o700); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(dst, []byte("hello, world!"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := ntos.CopyFile(dst, src, 0o600); err == nil {
		t.Error("CopyFile succeeded copying a directory to a file")
	}
}

func TestCopyFile_empty(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "ntgo-os-copy-file_empty")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	dst := filepath.Join(dir, "destination.txt")
	if err := ntos.CopyFile(dst, "", 0o600); err == nil {
		t.Error("CopyFile succeeded without source")
	}

	src := filepath.Join(dir, "source.txt")
	if err := os.WriteFile(src, []byte("hello, world!"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := ntos.CopyFile("", src, 0o600); err == nil {
		t.Error("CopyFile succeeded without destination")
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

	sbytes, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("can't read %s: %w", src, err)
	}

	dbytes, err := os.ReadFile(dst)
	if err != nil {
		return fmt.Errorf("can't read %s: %w", dst, err)
	}

	if !bytes.Equal(sbytes, dbytes) {
		return errors.New("files have different contents")
	}

	return nil
}
