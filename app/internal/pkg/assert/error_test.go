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
	"io"
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
)

// UT: Compare 2 (error) values for equality.
func TestErrorf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got %v, want %v."

	for tcName, tc := range map[string]struct {
		gotInput  error
		wantInput error
		want      string
	}{
		"When 'got' and 'want' are equal, the assertion should NOT fail.": {
			gotInput:  io.EOF,
			wantInput: io.EOF,
			want:      "",
		},
		"When 'got' and 'want' are <nil>, the assertion should NOT fail.": {
			gotInput:  nil,
			wantInput: nil,
			want:      "",
		},
		"When 'got' is <nil> and 'want' is NOT <nil>, the assertion should fail.": {
			gotInput:  nil,
			wantInput: io.EOF,
			want:      "Not equal - got <nil>, want EOF.",
		},
		"When 'got' is NOT <nil> and 'want' is <nil>, the assertion should fail.": {
			gotInput:  io.EOF,
			wantInput: nil,
			want:      "Not equal - got EOF, want <nil>.",
		},
		"When 'got' and want are NOT equal, the assertion should fail.": {
			gotInput:  io.ErrUnexpectedEOF,
			wantInput: io.EOF,
			want:      "Not equal - got unexpected EOF, want EOF.",
		},
	} {
		tcName, tc := tcName, tc // Rebind: Needed for parallel support.

		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Arrange.
			tbSpy := SpyTb(t)

			// Act.
			assert.Errorf(tbSpy, tc.gotInput, tc.wantInput, msgFmt, tc.gotInput, tc.wantInput)

			// Assert.
			if tbSpy.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, tc.want)
			}
		})
	}
}
