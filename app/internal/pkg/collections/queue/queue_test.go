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

// QA: Verify and measure the performance of the public API of the "queue" package.
package queue_test

import (
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
	"github.com/kdeconinck/align/internal/pkg/collections/queue"
)

// UT: Create a new [queue.Queue].
func TestNew(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	q := queue.New[int]()

	// Act.
	got, want := q.Len(), 0

	// Assert.
	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When creating a new 'Queue', it contains NO elements.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want, got)
}

// UT: Create a new [queue.Queue] with a specific capacity.
func TestWithCapacity(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	q := queue.WithCapacity[int](10)

	// Act.
	got, want := q.Len(), 0

	// Assert.
	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When creating a new 'Queue', it contains NO elements.\n"+
		"\033[32mExpected: %d.\033[0m\n"+
		"\033[31mActual:   %d.\033[0m\n\n", want, got)
}

// UT: Enqueue elements into a [queue.Queue].
func TestQueue_Enqueue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When enqueuing a single element, the total amount of elements is increased.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		q := queue.New[int]()

		// Act.
		q.Enqueue(1)

		got, want := q.Len(), 1

		// Assert.
		assert.Equalf(t, want, got, "\n\n"+
			"UT Name:  When enqueuing a single element, the total amount of elements is increased.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When enqueuing multiple elements, the total amount of elements matches the number enqueued.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		q := queue.New[int]()

		// Act.
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		got, want := q.Len(), 3

		// Assert.
		assert.Equalf(t, want, got, "\n\n"+
			"UT Name:  When enqueuing multiple elements, the total amount of elements matches the number enqueued.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Dequeue elements from a [queue.Queue].
func TestQueue_Dequeue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When dequeuing from an empty queue, false is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		q := queue.New[int]()

		// Act.
		_, ok := q.Dequeue()

		// Assert.
		assert.Falsef(t, ok, "\n\n"+
			"UT Name:  When dequeuing from an empty queue, false is returned.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", ok)
	})

	t.Run("When dequeuing, true is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		q := queue.New[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		// Act.
		_, got := q.Dequeue()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  When dequeuing, true is returned.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("When dequeuing, the first element is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		q := queue.New[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		// Act.
		got, _ := q.Dequeue()

		// Assert.
		want := 1

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When dequeuing, the first element is returned.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When dequeuing, the element is removed.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		q := queue.New[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		// Act.
		q.Dequeue()

		// Assert.
		got, want := q.Len(), 2

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When dequeuing, the element is removed.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

var benchmarkOutput int // Output of the benchmark(s). Used to avoid compiler optimizations.

// Benchmark(s): Enqueue elements in a [queue.Queue] that has NO predefined capacity.
func BenchmarkQueue_1(b *testing.B)       { benchmarkQueue(1, b) }
func BenchmarkQueue_10(b *testing.B)      { benchmarkQueue(10, b) }
func BenchmarkQueue_100(b *testing.B)     { benchmarkQueue(100, b) }
func BenchmarkQueue_1000(b *testing.B)    { benchmarkQueue(1_000, b) }
func BenchmarkQueue_1000000(b *testing.B) { benchmarkQueue(1_000_000, b) }

// Benchmark(s): Enqueue elements in a [queue.Queue] that has enough capacity for the items.
func BenchmarkPreallocQueue_1(b *testing.B)       { benchmarkPreallocQueue(1, b) }
func BenchmarkPreallocQueue_10(b *testing.B)      { benchmarkPreallocQueue(10, b) }
func BenchmarkPreallocQueue_100(b *testing.B)     { benchmarkPreallocQueue(100, b) }
func BenchmarkPreallocQueue_1000(b *testing.B)    { benchmarkPreallocQueue(1_000, b) }
func BenchmarkPreallocQueue_1000000(b *testing.B) { benchmarkPreallocQueue(1_000_000, b) }

// Benchmark: Measure the performance of enqueuing elements in a [queue.Queue] that has NO capacity.
// Parameters:
// - count: The amount of elements to enqueue.
// - b:     The [testing.B] instance.
func benchmarkQueue(count int, b *testing.B) {
	var q queue.Queue[int]

	for b.Loop() {
		q := queue.New[int]()

		for range count {
			q.Enqueue(0)
		}
	}

	benchmarkOutput = q.Len()
}

// Benchmark: Measure the performance of enqueuing elements in a [queue.Queue] that has enough capacity for the items.
// Parameters:
// - count: The amount of elements to enqueue.
// - b:     The [testing.B] instance.
func benchmarkPreallocQueue(count int, b *testing.B) {
	// Arrange.
	var q queue.Queue[int]

	for b.Loop() {
		q := queue.WithCapacity[int](count)

		for range count {
			q.Enqueue(0)
		}
	}

	benchmarkOutput = q.Len()
}
