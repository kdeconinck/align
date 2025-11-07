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

// UT: Compare 2 values for equality.
func TestEqualf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got %t, want %t."

	for tcName, tc := range map[string]struct {
		gotInput  bool
		wantInput bool
		want      string
	}{
		"When 'got' and 'want' are equal, the assertion should NOT fail.": {
			gotInput:  true,
			wantInput: true,
			want:      "",
		},
		"When 'got' and 'want' are NOT equal, the assertion should fail.": {
			gotInput:  true,
			wantInput: false,
			want:      "Not equal - got true, want false.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.Equalf(tbSpy, tc.gotInput, tc.wantInput, msgFmt, tc.gotInput, tc.wantInput)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}

// UT: Compare 2 slices for equality.
func TestEqualSf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got %v, want %v."

	for tcName, tc := range map[string]struct {
		gotInput  []int
		wantInput []int
		want      string
	}{
		"When 'got' and 'want' are equal, the assertion should NOT fail.": {
			gotInput:  newSlice(1, 2, 3),
			wantInput: newSlice(1, 2, 3),
			want:      "",
		},
		"When 'got' and 'want' have a different length, the assertion should fail.": {
			gotInput:  newSlice(1, 2, 3, 4, 5),
			wantInput: newSlice(1, 2, 3),
			want:      "Not equal - got [1 2 3 4 5], want [1 2 3].",
		},
		"When 'got' and 'want' have a different order, the assertion should fail.": {
			gotInput:  newSlice(1, 2, 3),
			wantInput: newSlice(3, 2, 1),
			want:      "Not equal - got [1 2 3], want [3 2 1].",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.EqualSf(tbSpy, tc.gotInput, tc.wantInput, msgFmt, tc.gotInput, tc.wantInput)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}
