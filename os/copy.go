// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package os

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Copy copies src content into dst. If dst doesn't exists, it will be created.
// src mode will override dst mode. Only file-file or directory-directory
// operations should be performed.
func Copy(dst, src string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return NewCopyError(src, dst, "cannot stat source file", err)
	}

	if sfi.IsDir() {
		return CopyDir(dst, src, sfi.Mode())
	}

	return CopyFile(dst, src, sfi.Mode())
}

// CopyDir copies src content recursively into dst, if dst doesn't exists it
// will be created. mode will be the new dst mode, its content will preserve
// the origin mode.
func CopyDir(dst, src string, mode os.FileMode) error {
	if err := isInside(dst, src); err != nil {
		return err
	}

	fn := func(srcpath string, fi os.FileInfo, err error) error {
		if err != nil {
			return NewCopyError(src, dst, "cannot stat "+srcpath, err)
		}

		dest := filepath.Clean(strings.Replace(srcpath, src, dst, 1))

		if fi.IsDir() {
			err := os.Mkdir(dest, fi.Mode())
			if err != nil && !errors.Is(err, os.ErrExist) {
				return NewCopyError(src, dst, "cannot create directory "+dest, err)
			}

			return nil
		}

		return CopyFile(dest, srcpath, fi.Mode())
	}

	return filepath.Walk(src, fn)
}

// CopyFile copies src content into dst, if dst exists it will be truncated.
// mode will be the new dst mode.
func CopyFile(dst, src string, mode os.FileMode) error {
	from, err := os.Open(src)
	if err != nil {
		return NewCopyError(src, dst, "cannot open source file", err)
	}

	defer from.Close()

	to, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return NewCopyError(src, dst, "cannot open destination file", err)
	}

	defer to.Close()

	if _, err = io.Copy(to, from); err != nil {
		return NewCopyError(src, dst, "cannot copy data", err)
	}

	return nil
}

// CopyError records an error during a copy operation. If Err is nil, it means
// the error doesn't wrap any error from another package.
type CopyError struct {
	Src, Dst string
	Reason   string
	Err      error
}

// NewCopyError creates a CopyError.
func NewCopyError(src, dst, msg string, err error) error {
	return &CopyError{
		Src:    src,
		Dst:    dst,
		Reason: msg,
		Err:    err,
	}
}

// Error implements the error interface.
func (e *CopyError) Error() string {
	err := "cannot copy " + e.Src + " to " + e.Dst + ", " + e.Reason

	if e.Err != nil {
		err += ": " + e.Err.Error()
	}

	return err
}

// Unwrap allows to use functions from errors package over CopyError.
func (e *CopyError) Unwrap() error {
	return e.Err
}

func isInside(dst, src string) error {
	srcabs, err := filepath.Abs(src)
	if err != nil {
		return NewCopyError(src, dst, "cannot get source absolute path", err)
	}

	dstabs, err := filepath.Abs(dst)
	if err != nil {
		return NewCopyError(src, dst, "cannot get destination absolute path", err)
	}

	if srcabs[len(srcabs)-1] != filepath.Separator {
		srcabs += string(filepath.Separator)
	}

	if strings.HasPrefix(dstabs, srcabs) {
		return NewCopyError(src, dst, "cannot copy a directory into itself", nil)
	}

	return nil
}
