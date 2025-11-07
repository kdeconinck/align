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

// UT: Compare the length of a slice against 0.
func TestEmptyf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal: # Of elements - got %d, want 0."

	for tcName, tc := range map[string]struct {
		gotInput []int
		want     string
	}{
		"When 'got' has NO elements, the assertion should NOT fail.": {
			gotInput: newSlice(),
			want:     "",
		},
		"When 'got' has elements, the assertion should fail.": {
			gotInput: newSlice(1, 2, 3, 4, 5),
			want:     "Not equal: # Of elements - got 5, want 0.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.IsEmptyf(tbSpy, tc.gotInput, msgFmt, len(tc.gotInput))

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}
