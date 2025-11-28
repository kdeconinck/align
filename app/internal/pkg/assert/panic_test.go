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
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
)

// UT: Verify if a function caused a panic.
func TestPanicf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		fn   func(int)
		arg  int
		want string
	}{
		"When the function did NOT cause a panic, the assertion should fail.": {
			fn:   func(_ int) {},
			arg:  1,
			want: "1 - The function did NOT panic.",
		},
		"When the function did cause a panic, the assertion should NOT fail.": {
			fn: func(_ int) {
				panic("unhandled")
			},
			arg:  1,
			want: "",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.Panicf(tbSpy, func() { tc.fn(tc.arg) }, "%d - The function did NOT panic.", tc.arg)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}
