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
	"unsafe"

	"github.com/kdeconinck/align/internal/pkg/assert"
)

// UT: Compare a value against <nil>.
func TestNilf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	tbSpy := SpyTb(t)

	// Act.
	assert.Nilf(tbSpy, 1, "Not equal - got %v, want <nil>.", 1)

	// Assert.
	want := "Not equal - got 1, want <nil>."

	if tbSpy.failureMsg != want {
		t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
	}
}

// UT: Compare a value against NOT <nil>.
func TestNotNilf(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	const msgFmt string = "Not equal - got default '%T', want NOT <nil>."

	t.Run("When 'got' is NOT <nil>, the assertion should NOT fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		assert.NotNilf(tbSpy, 1, msgFmt, 1)

		// Assert.
		if tbSpy.failureMsg != "" {
			t.Fatalf("Failure message = \"%v\", want \"\"", tbSpy.failureMsg)
		}
	})

	t.Run("When 'got' is the default 'any' value, the assertion should fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		var value any

		assert.NotNilf(tbSpy, value, msgFmt, value)

		// Assert.
		want := "Not equal - got default '<nil>', want NOT <nil>."

		if tbSpy.failureMsg != want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
		}
	})

	t.Run("When 'got' is the default 'chan' value, the assertion should fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		var value chan int

		assert.NotNilf(tbSpy, value, msgFmt, value)

		// Assert.
		want := "Not equal - got default 'chan int', want NOT <nil>."

		if tbSpy.failureMsg != want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
		}
	})

	t.Run("When 'got' is the default 'func' value, the assertion should fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		var value func(int) bool

		assert.NotNilf(tbSpy, value, msgFmt, value)

		// Assert.
		want := "Not equal - got default 'func(int) bool', want NOT <nil>."

		if tbSpy.failureMsg != want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
		}
	})

	t.Run("When 'got' is the default 'map' value, the assertion should fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		var value map[int]bool

		assert.NotNilf(tbSpy, value, msgFmt, value)

		// Assert.
		want := "Not equal - got default 'map[int]bool', want NOT <nil>."

		if tbSpy.failureMsg != want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
		}
	})

	t.Run("When 'got' is the default 'pointer' value, the assertion should fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		var value *int

		assert.NotNilf(tbSpy, value, msgFmt, value)

		// Assert.
		want := "Not equal - got default '*int', want NOT <nil>."

		if tbSpy.failureMsg != want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
		}
	})

	t.Run("When 'got' is the default 'slice' value, the assertion should fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		var value []int

		assert.NotNilf(tbSpy, value, msgFmt, value)

		// Assert.
		want := "Not equal - got default '[]int', want NOT <nil>."

		if tbSpy.failureMsg != want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
		}
	})

	t.Run("When 'got' is the default 'unsafe pointer' value, the assertion should fail.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		tbSpy := SpyTb(t)

		// Act.
		var value unsafe.Pointer

		assert.NotNilf(tbSpy, value, msgFmt, value)

		// Assert.
		want := "Not equal - got default 'unsafe.Pointer', want NOT <nil>."

		if tbSpy.failureMsg != want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", tbSpy.failureMsg, want)
		}
	})
}
