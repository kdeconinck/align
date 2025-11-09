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

// Package mvmap provided a tiny, generic, append-friendly multi-value map.
package mvmap

// MvMap maps a key K to zero or more values of type V.
type MvMap[K comparable, V any] struct {
	data map[K][]V
}

// New returns a new [MvMap] mapping K to []V.
func New[K comparable, V any]() *MvMap[K, V] {
	return &MvMap[K, V]{
		data: make(map[K][]V),
	}
}

// WithCapacity returns an empty [Set] that has capacity of cap.
func WithCapacity[K comparable, V any](cap int) *MvMap[K, V] {
	return &MvMap[K, V]{
		data: make(map[K][]V, cap),
	}
}

// SetKeyCap allocates a slice with a capacity of size for k.
func (m *MvMap[K, V]) SetKeyCap(k K, size int) {
	m.data[k] = make([]V, 0, size)
}

// Put stores v as a value of k in the map.
func (m *MvMap[K, V]) Put(k K, v V) {
	m.data[k] = append(m.data[k], v)
}

// Get returns the values of k in the map.
func (m *MvMap[K, V]) Get(k K) []V {
	return m.data[k]
}

// Len returns the total amount of keys in the map.
func (m *MvMap[K, V]) Len() int {
	return len(m.data)
}

// Keys returns all distinct keys in the map.
func (m *MvMap[K, V]) Keys() []K {
	out := make([]K, 0, m.Len())

	for v := range m.data {
		out = append(out, v)
	}

	return out
}
