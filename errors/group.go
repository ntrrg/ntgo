// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors

// Group groups all given errors into a single error. The returned group is
// compatible with errors.Is. This does not wrap errors, only creates a useful
// container for them.
func Group(errs ...error) error {
	return &group{errs: errs}
}

// Split separates a error group. If err is not a group, a single element slice
// containing it will be returned.
func Split(errs error) []error {
	g, ok := errs.(*group) // nolint:errorlint
	if !ok {
		return []error{errs}
	}

	return g.errs
}

type group struct {
	errs []error
}

func (g *group) Error() string {
	err := ""
	l := len(g.errs)

	for i, e := range g.errs {
		err += "* " + e.Error()

		if i < l-1 {
			err += "; "
		}
	}

	return err
}

func (g *group) Is(target error) bool {
	return Any(target, g.errs...)
}
