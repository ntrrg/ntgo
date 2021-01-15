// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package runtime_test

import (
	"runtime"
	"testing"

	ntruntime "go.ntrrg.dev/ntgo/runtime"
)

func TestIsBigEndian(t *testing.T) {
	t.Parallel()

	var want bool

	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "s390x":
		want = true
	default:
		want = false
	}

	if ntruntime.IsBigEndian() != want {
		t.Errorf("IsBigEndian() returns false in a big-endian system")
	}
}

func TestIsLittleEndian(t *testing.T) {
	t.Parallel()

	var want bool

	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "s390x":
		want = false
	default:
		want = true
	}

	if ntruntime.IsLittleEndian() != want {
		t.Errorf("IsLittleEndian() returns false in a little-endian system")
	}
}

func TestWordSize(t *testing.T) {
	t.Parallel()

	var want uintptr

	switch {
	case runtime.GOARCH == "386":
		want = 4
	default:
		want = 8
	}

	if got := ntruntime.WordSize(); got != want {
		t.Errorf(
			"WordSize() returns invalid word size. got: %v, want: %v",
			got, want,
		)
	}
}
