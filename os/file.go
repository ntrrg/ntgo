// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package os

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Cp behaves like the Unix command cp with -rf flags. The possible scenarios
// are:
//
// * [FAIL] No source.
//
// * [FAIL] Non existent source.
//
// * [FAIL] Multiple non existent elements in destination path.
//
// * [FAIL] Bad permissions.
//
// * [PASS] File to non existent file, dst will be created with the same
// content as src.
//
// * [PASS] File to file, dst will be overwritten by src.
//
// * [PASS] File to directory, dst will be a file inside dst with the same base
// name as src.
//
// * [FAIL] Directory to file.
//
// * [PASS] Directory to non existent directory, dst will be created with the
// same content as src.
//
// * [PASS] Directory to directory, dst will be a directory inside dst with the
// same base name as src.
//
// * [FAIL] Directory to itself.
//
// * [FAIL] Multiple elements to file.
//
// * [FAIL] Multiple elements to non existent directory.
//
// * [PASS] Multiple elements to directory, any src elements will be
// created/overwritten in dst.
func Cp(dst string, src ...string) error {
	dfi, err := os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	n := len(src)

	switch {
	case n == 0:
		return errors.New("missing source")
	case n == 1 && dfi != nil && dfi.IsDir():
		dst = filepath.Join(dst, filepath.Base(src[0]))
	case n > 1 && dfi == nil:
		return fmt.Errorf("%v doesn't exists", dst)
	case n > 1 && !dfi.IsDir():
		return fmt.Errorf("%v is not a directory", dst)
	}

	var errs []string
	var fn func(string, string, os.FileInfo) error

	for _, entry := range src {
		sfi, err := os.Stat(entry)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}

		if sfi.IsDir() {
			fn = copyDir
		} else {
			fn = copyFile
		}

		dst := dst
		if n > 1 {
			dst = filepath.Join(dst, sfi.Name())
		}

		if err := fn(dst, entry, sfi); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}

func copyDir(dst, src string, sfi os.FileInfo) error {
	srcabs, err := filepath.Abs(src)
	if err != nil {
		return err
	}

	dstabs, err := filepath.Abs(dst)
	if err != nil {
		return err
	}

	if strings.HasPrefix(dstabs, srcabs) {
		return fmt.Errorf("cannot copy directory %v into itself (%s)", src, dst)
	}

	if err := os.MkdirAll(dst, sfi.Mode()); err != nil {
		return err
	}

	var errs []string

	fn := func(srcpath string, fi os.FileInfo, err error) error {
		if src == srcpath {
			return nil
		}

		if err != nil {
			errs = append(errs, err.Error())
			return nil
		}

		dst := filepath.Clean(strings.Replace(srcpath, src, dst, 1))

		if fi.IsDir() {
			if err := os.MkdirAll(dst, fi.Mode()); err != nil {
				errs = append(errs, err.Error())
				return nil
			}
		} else {
			if err := copyFile(dst, srcpath, fi); err != nil {
				errs = append(errs, err.Error())
				return nil
			}
		}

		return nil
	}

	if err := filepath.Walk(src, fn); err != nil {
		return err
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}

func copyFile(dst, src string, sfi os.FileInfo) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}

	defer from.Close()

	to, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, sfi.Mode())
	if err != nil {
		return err
	}

	defer to.Close()

	if _, err = io.Copy(to, from); err != nil {
		return err
	}

	return nil
}
