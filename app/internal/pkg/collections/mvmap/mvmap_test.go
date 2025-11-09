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

// QA: Verify and measure the performance of the public API of the "mvmap" package.
package mvmap_test

import (
	"sort"
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
	"github.com/kdeconinck/align/internal/pkg/collections/mvmap"
)

// UT: Create a new [mvmap.MvMap].
func TestNew(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	m := mvmap.New[int, int]()

	// Act.
	got, want := m.Len(), 0

	// Assert.
	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When creating a new 'MvMap', it contains NO elements.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want, got)
}

// UT: Create a new [mvmap.MvMap] with a specific capacity.
func TestWithCapacity(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	m := mvmap.WithCapacity[int, int](10)

	// Act.
	got, want := m.Len(), 0

	// Assert.
	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When creating a new 'MvMap', it contains NO elements.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want, got)
}

// UT: Set a specific capacity for an [mvmap.MvMap] key.
func TestMvMap_SetKeyCap(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	m := mvmap.New[int, int]()

	// Act.
	m.SetKeyCap(0, 100)

	// Assert.
	got, want := m.Get(0), newSlice()

	assert.EqualSf(t, got, want, "\n\n"+
		"UT Name:  When specifying the size of a key inside a 'MvMap', the key contains NO elements.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want, got)
}

// UT: Put elements into a [mvmap.MvMap].
func TestMvMap_Put(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When adding a single element, the total amount of elements is increased.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()

		// Act.
		m.Put(1, 10)

		got, want := m.Len(), 1

		// Assert.
		assert.Equalf(t, want, got, "\n\n"+
			"UT Name:  When adding a single element, the total amount of elements is increased.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When adding multiple elements (with shared keys), the total amount of elements matches the number of unique keys.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()

		// Act.
		m.Put(1, 10)
		m.Put(1, 100)
		m.Put(2, 20)

		got, want := m.Len(), 2

		// Assert.
		assert.Equalf(t, want, got, "\n\n"+
			"UT Name:  When adding multiple elements (with shared keys), the total amount of elements matches the number of unique keys.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Get elements into a [mvmap.MvMap].
func TestMvMap_Get(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When requesting the values of a key that's NOT in the 'MvMap', an empty slice is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()

		// Act.
		got, want := m.Get(1), newSlice()

		// Assert.
		assert.EqualSf(t, want, got, "\n\n"+
			"UT Name:  When requesting the values of a key that's NOT in the 'MvMap', an empty slice is returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When requesting the values of a key that's in the 'MvMap', the values of the key are returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()
		m.Put(1, 10)
		m.Put(1, 100)
		m.Put(2, 20)

		// Act.
		got, want := m.Get(1), newSlice(10, 100)

		// Assert.
		assert.EqualSf(t, want, got, "\n\n"+
			"UT Name:  When requesting the values of a key that's in the 'MvMap', the values of the key are returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Count the elements of a [mvmap.MvMap].
func TestMvMap_Len(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When there are NO elements, 0 is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()

		// Act.
		got, want := m.Len(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When there are NO elements, 0 is returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When there are elements (with shared keys), the total amount of elements matches the number of unique keys.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()
		m.Put(1, 10)
		m.Put(1, 100)
		m.Put(2, 20)

		// Act.
		got, want := m.Len(), 2

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When there are elements (with shared keys), the total amount of elements matches the number of unique keys.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Get the keys of a [mvmap.MvMap].
func TestMvMap_Keys(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When there are NO elements, an empty slice is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()

		// Act.
		got, want := m.Keys(), newSlice()

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  When there are NO elements, 0 is returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When there are elements (with shared keys), the unique keys are returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		m := mvmap.New[int, int]()
		m.Put(1, 10)
		m.Put(1, 100)
		m.Put(2, 20)

		// Act.
		got, want := m.Keys(), newSlice(1, 2)

		// Assert.
		sort.Ints(got)

		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  When there are elements (with shared keys), the unique keys are returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// Utility: Return a slice of integers, containing args.
func newSlice(args ...int) []int {
	container := make([]int, len(args))
	copy(container, args)

	return container
}

var benchmarkOutput int // Output of the benchmark(s). Used to avoid compiler optimizations.

// Benchmark(s): Add keys in a [mvmap.MvMap] that has NO predefined capacity.
func BenchmarkMap_AddKey_1(b *testing.B)       { benchmarkMap_AddKey(1, b) }
func BenchmarkMap_AddKey_10(b *testing.B)      { benchmarkMap_AddKey(10, b) }
func BenchmarkMap_AddKey_100(b *testing.B)     { benchmarkMap_AddKey(100, b) }
func BenchmarkMap_AddKey_1000(b *testing.B)    { benchmarkMap_AddKey(1_000, b) }
func BenchmarkMap_AddKey_1000000(b *testing.B) { benchmarkMap_AddKey(1_000_000, b) }

// Benchmark(s): Add keys in a [mvmap.MvMap] that has enough capacity for the keys.
func BenchmarkPreallocMap_AddKey_1(b *testing.B)       { benchmarkPreallocMap_AddKey(1, b) }
func BenchmarkPreallocMap_AddKey_10(b *testing.B)      { benchmarkPreallocMap_AddKey(10, b) }
func BenchmarkPreallocMap_AddKey_100(b *testing.B)     { benchmarkPreallocMap_AddKey(100, b) }
func BenchmarkPreallocMap_AddKey_1000(b *testing.B)    { benchmarkPreallocMap_AddKey(1_000, b) }
func BenchmarkPreallocMap_AddKey_1000000(b *testing.B) { benchmarkPreallocMap_AddKey(1_000_000, b) }

// Benchmark(s): Add values (for a single key) in a [mvmap.MvMap] that has enough capacity for the values.
func BenchmarkPreallocKey_AddValue_1(b *testing.B)       { benchmarkPreallocKey_AddValue(1, b) }
func BenchmarkPreallocKey_AddValue_10(b *testing.B)      { benchmarkPreallocKey_AddValue(10, b) }
func BenchmarkPreallocKey_AddValue_100(b *testing.B)     { benchmarkPreallocKey_AddValue(100, b) }
func BenchmarkPreallocKey_AddValue_1000(b *testing.B)    { benchmarkPreallocKey_AddValue(1_000, b) }
func BenchmarkPreallocKey_AddValue_1000000(b *testing.B) { benchmarkPreallocKey_AddValue(1_000_000, b) }

// Benchmark: Measure the performance of enqueuing elements in a [mvmap.MvMap] that has NO capacity.
// Parameters:
// - count: The amount of keys to add.
// - b:     The [testing.B] instance.
func benchmarkMap_AddKey(count int, b *testing.B) {
	// Arrange.
	var m mvmap.MvMap[int, int]

	for b.Loop() {
		m := mvmap.New[int, int]()

		for idx := 0; idx < count; idx++ {
			m.Put(idx, 0)
		}
	}

	benchmarkOutput = m.Len()
}

// Benchmark: Measure the performance of adding elements in a [mvmap.MvMap] that has enough capacity for the keys.
// Parameters:
// - count: The amount of keys to add.
// - b:     The [testing.B] instance.
func benchmarkPreallocMap_AddKey(count int, b *testing.B) {
	// Arrange.
	var m mvmap.MvMap[int, int]

	for b.Loop() {
		m := mvmap.WithCapacity[int, int](count)

		for idx := 0; idx < count; idx++ {
			m.Put(idx, 0)
		}
	}

	benchmarkOutput = m.Len()
}

// Benchmark: Measure the performance of adding values (for a single key) in a [mvmap.MvMap] that has enough capacity.
// Parameters:
// - count: The amount of values to add.
// - b:     The [testing.B] instance.
func benchmarkPreallocKey_AddValue(count int, b *testing.B) {
	// Arrange.
	var m mvmap.MvMap[int, int]

	for b.Loop() {
		m := mvmap.New[int, int]()
		m.SetKeyCap(0, count)

		for idx := 0; idx < count; idx++ {
			m.Put(0, idx)
		}
	}

	benchmarkOutput = m.Len()
}
