// =====================================================================================================================
// = LICENSE:       Copyright (c) 2025 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package assert defines functions for making assertions in Go's standard testing package.
package assert

import (
	"errors"
	"testing"
)

// Errorf compares got against want for equality.
// If they aren't equal, tb is marked as failed and the execution is terminated.
// The failure message is formatted using [testing.TB.Fatalf] using format and args.
func Errorf(tb testing.TB, got, want error, format string, args ...any) {
	tb.Helper()

	switch {
	case got == nil && want == nil:
		return

	case got == nil && want != nil:
		tb.Fatalf(format, args...)

	case got != nil && want == nil:
		tb.Fatalf(format, args...)

	default:
		if !errors.Is(got, want) {
			tb.Fatalf(format, args...)
		}
	}
}
