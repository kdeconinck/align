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

// Package set implements a tiny, generic set (unique elements) implementation built on top of Go's built-in map.
package set

// Set is a container holding unique T elements.
type Set[T comparable] map[T]struct {
	// NOTE: Intentionally left blank.
}

// New returns an empty [Set].
func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// WithCapacity returns an empty [Set] that has capacity of cap.
func WithCapacity[T comparable](cap int) Set[T] {
	return make(map[T]struct{}, cap)
}

// Add adds v into the set.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Has reports whether v is present in the set.
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]

	return ok
}

// Values returns a slice containing all elements currently in the set.
func (s Set[T]) Values() []T {
	out := make([]T, 0, len(s))

	for v := range s {
		out = append(out, v)
	}

	return out
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}
