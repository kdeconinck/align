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

// Package queue implements a tiny, generic FIFO (first in, first out) queue.
package queue

// Queue is a FIFO (first in, first out) queue, containing elements of type T.
type Queue[T any] struct {
	data []T
}

// New returns an empty [Queue].
func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// WithCapacity returns an empty [Queue] that has capacity of cap.
func WithCapacity[T any](cap int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0, cap),
	}
}

// Enqueue pushes v at the end of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue pops from the front of the queue.
// It returns the popped element and true if an element was popped.
// If the queue contains no elements to pop, the second return value is false.
func (q *Queue[T]) Dequeue() (v T, ok bool) {
	if len(q.data) == 0 {
		return v, false
	}

	v = q.data[0]
	q.data = q.data[1:]

	return v, true
}

// Len returns the amount of items in the queue.
func (q *Queue[T]) Len() int {
	return len(q.data)
}
