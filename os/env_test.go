// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package os_test

import (
	"os"
	"testing"

	ntos "go.ntrrg.dev/ntgo/os"
)

func TestGetenvOr(t *testing.T) {
	t.Parallel()

	cases := []struct {
		label string
		key   string
		set   string
		val   string
		want  string
	}{
		{
			label: "Existing",
			key:   "X_GETENVOR",
			set:   "X",
			want:  "X",
		},

		{
			label: "Existing with default value",
			key:   "X_GETENVOR_DEFAULT",
			set:   "X",
			val:   "Y",
			want:  "X",
		},

		{
			label: "Missing",
			key:   "X_GETENVOR_MISSING",
			want:  "",
		},

		{
			label: "Missing with default value",
			key:   "X_GETENVOR_MISSING_DEFAULT",
			val:   "Y",
			want:  "Y",
		},
	}

	for _, c := range cases {
		if c.set != "" {
			os.Setenv(c.key, c.set)
		}

		got := ntos.GetenvOr(c.key, c.val)
		if got != c.want {
			t.Errorf("[%s] invalid value. got: %q; want: %q", c.label, got, c.want)
		}
	}
}
