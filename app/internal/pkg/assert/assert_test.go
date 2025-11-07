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

// QA: Verify and measure the performance of the public API of the "assert" package.
package assert_test

import (
	"fmt"
	"testing"
)

// TbSpy is a test double (Spy) that wraps a [testing.TB] to record failure messages instead of stopping the test.
type TbSpy struct {
	testing.TB        // The wrapped [testing.TB].
	failureMsg string // The failure message, if any.
}

// SpyTb creates and initializes a new [TbSpy].
// It wraps the provided [testing.TB] instance for internal use by methods that delegate to the real object.
func SpyTb(tb testing.TB) *TbSpy {
	return &TbSpy{
		TB: tb,
	}
}

// Fatalf formats using [fmt.Sprintf] using format and args and stores the result.
func (tb *TbSpy) Fatalf(format string, args ...any) {
	tb.failureMsg = fmt.Sprintf(format, args...)
}

// Utility: Build and return a slice of integers, containing args.
func newSlice(args ...int) []int {
	container := make([]int, len(args))
	copy(container, args)

	return container
}
