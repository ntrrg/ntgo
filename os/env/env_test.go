// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package env_test

import (
	"os"
	"testing"

	"go.ntrrg.dev/ntgo/os/env"
)

func TestGetOr(t *testing.T) {
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
			key:   "X_TEST_OS_ENV_GETOR",
			set:   "X",
			want:  "X",
		},

		{
			label: "Existing with default value",
			key:   "X_TEST_OS_ENV_GETOR_DEFAULT",
			set:   "X",
			val:   "Y",
			want:  "X",
		},

		{
			label: "Missing",
			key:   "X_TEST_OS_ENV_GETOR_MISSING",
			want:  "",
		},

		{
			label: "Missing with default value",
			key:   "X_TEST_OS_ENV_GETOR_MISSING_DEFAULT",
			val:   "Y",
			want:  "Y",
		},
	}

	for _, c := range cases {
		if c.set != "" {
			os.Setenv(c.key, c.set)
		}

		got, err := env.GetOr(c.key, c.val, env.String)
		if err != nil {
			t.Errorf("[%s] can't get value: %v", c.label, err)
			continue
		}

		if got != c.want {
			t.Errorf("[%s] invalid value. got: %q; want: %q", c.label, got, c.want)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	key := "X_BENCH_OS_ENV_GET"
	val := "value"
	want := val

	b.Setenv(key, val)

	b.Run("Stdlib", func(b *testing.B) {
		var v string

		for i := 0; i < b.N; i++ {
			v = os.Getenv(key)
			if v != want {
				b.Fatalf("invalid value. got: %q, want: %q", v, want)
			}
		}
	})

	b.Run("Get", func(b *testing.B) {
		var v string

		for i := 0; i < b.N; i++ {
			v, _ = env.Get(key, env.String) //nolint:errcheck
			if v != want {
				b.Fatalf("invalid value. got: %q, want: %q", v, want)
			}
		}
	})

	b.Run("GetOr", func(b *testing.B) {
		var v string

		for i := 0; i < b.N; i++ {
			v, _ = env.GetOr(key, want, env.String) //nolint:errcheck
			if v != want {
				b.Fatalf("invalid value. got: %q, want: %q", v, want)
			}
		}
	})

	b.Run("GetOr_Unset", func(b *testing.B) {
		var v string

		for i := 0; i < b.N; i++ {
			v, _ = env.GetOr("X_BENCH_OS_ENV_GET_UNSET", want, env.String) //nolint:errcheck
			if v != want {
				b.Fatalf("invalid value. got: %q, want: %q", v, want)
			}
		}
	})
}
