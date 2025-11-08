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

// QA: Verify and measure the performance of the public API of the "set" package.
package set_test

import (
	"sort"
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
	"github.com/kdeconinck/align/internal/pkg/collections/set"
)

// UT: Create a new [set.Set].
func TestNew(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	s := set.New[int]()

	// Act.
	got, want := s.Len(), 0

	// Assert.
	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When creating a new 'Set', it contains NO elements.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want, got)
}

// UT: Create a new [set.Set] with a specific capacity.
func TestWithCapacity(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	s := set.WithCapacity[int](10)

	// Act.
	got, want := s.Len(), 0

	// Assert.
	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When creating a new 'Set', it contains NO elements.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want, got)
}

// UT: Add a new element to a [set.Set].
func TestSet_Add(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When adding a new element, the total amount of elements is increased.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := set.New[int]()

		// Act.
		s.Add(1)

		got, want := s.Len(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When adding a new element, the total amount of elements is increased.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When adding a new element, the element is present.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := set.New[int]()

		// Act.
		s.Add(1)

		// Assert.
		got := s.Has(1)

		assert.Truef(t, got, "\n\n"+
			"UT Name:  When adding a new element, the element is present.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("When adding an element multiple times, the total amount of elements is increased only once.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := set.New[int]()

		// Act.
		s.Add(1)
		s.Add(1)

		// Assert.
		got, want := s.Len(), 1

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When adding an element multiple times, the total amount of elements is increased only once\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When adding an element multiple times, the element is present.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := set.New[int]()

		// Act.
		s.Add(1)
		s.Add(1)

		// Assert.
		got := s.Has(1)

		assert.Truef(t, got, "\n\n"+
			"UT Name:  When adding a new element multiple times, the element is present.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Verify is an element is present in a [set.Set].
func TestSet_Has(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When the element is NOT present, false is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := set.New[int]()
		s.Add(1)
		s.Add(3)

		// Act.
		got := s.Has(2)

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  When the element is NOT present, false is returned.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("When the element is present, true is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := set.New[int]()
		s.Add(1)
		s.Add(2)
		s.Add(3)

		// Act.
		got := s.Has(2)

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  When the element is present, true is returned.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Return all elements from a [set.Set].
func TestSet_Values(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When requesting all the elements, all the elements are returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := set.New[int]()
		s.Add(1)
		s.Add(2)
		s.Add(3)

		// Act.
		got, want := s.Values(), newSlice(1, 2, 3)

		// Assert.
		sort.Ints(got)

		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  When requesting all the elements, all the elements are returned.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})
}

var benchmarkOutput int // Output of the benchmark(s). Used to avoid compiler optimizations.

// Benchmark(s): Add elements in a [set.Set] that has NO predefined capacity.
func BenchmarkSet_1(b *testing.B)       { benchmarkSet(1, b) }
func BenchmarkSet_10(b *testing.B)      { benchmarkSet(10, b) }
func BenchmarkSet_100(b *testing.B)     { benchmarkSet(100, b) }
func BenchmarkSet_1000(b *testing.B)    { benchmarkSet(1_000, b) }
func BenchmarkSet_1000000(b *testing.B) { benchmarkSet(1_000_000, b) }

// Benchmark(s): Add elements in a [set.Set] that has enough capacity for the items.
func BenchmarkPreallocSet_1(b *testing.B)       { benchmarkPreallocSet(1, b) }
func BenchmarkPreallocSet_10(b *testing.B)      { benchmarkPreallocSet(10, b) }
func BenchmarkPreallocSet_100(b *testing.B)     { benchmarkPreallocSet(100, b) }
func BenchmarkPreallocSet_1000(b *testing.B)    { benchmarkPreallocSet(1_000, b) }
func BenchmarkPreallocSet_1000000(b *testing.B) { benchmarkPreallocSet(1_000_000, b) }

// Benchmark: Measure the performance of adding elements in a [set.Set] that has NO capacity.
// Parameters:
// - count: The amount of elements to add.
// - b:     The [testing.B] instance.
func benchmarkSet(count int, b *testing.B) {
	var s set.Set[int]

	for b.Loop() {
		s := set.New[int]()

		for idx := 0; idx < count; idx++ {
			s.Add(idx)
		}
	}

	benchmarkOutput = s.Len()
}

// Benchmark: Measure the performance of adding elements in a [set.Set] that has enough capacity for the items.
// Parameters:
// - count: The amount of elements to set.
// - b:     The [testing.B] instance.
func benchmarkPreallocSet(count int, b *testing.B) {
	// Arrange.
	var s set.Set[int]

	for b.Loop() {
		s := set.WithCapacity[int](count)

		for idx := 0; idx < count; idx++ {
			s.Add(idx)
		}
	}

	benchmarkOutput = s.Len()
}

// Utility: Return a slice of integers, containing args.
func newSlice(args ...int) []int {
	container := make([]int, len(args))
	copy(container, args)

	return container
}
