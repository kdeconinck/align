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

// UT: Compare a value against true.
func TestTruef(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got %t, want true."

	for tcName, tc := range map[string]struct {
		gotInput bool
		want     string
	}{
		"When 'got' is true, the assertion should NOT fail.": {
			gotInput: true,
			want:     "",
		},
		"When 'got' is false, the assertion should fail.": {
			gotInput: false,
			want:     "Not equal - got false, want true.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.Truef(tbSpy, tc.gotInput, msgFmt, tc.gotInput)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}

// UT: Compare a value against false.
func TestFalsef(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "UT Failed: Bool value - got %t, want false."

	for tcName, tc := range map[string]struct {
		gotInput bool
		want     string
	}{
		"When 'got' is false, the assertion should NOT fail.": {
			gotInput: false,
			want:     "",
		},
		"When 'got' is true, the assertion should fail.": {
			gotInput: true,
			want:     "UT Failed: Bool value - got true, want false.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.Falsef(tbSpy, tc.gotInput, msgFmt, tc.gotInput)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}
