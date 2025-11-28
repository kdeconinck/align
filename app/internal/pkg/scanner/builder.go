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

// Package scanner implements a mechanism for transforming patterns into a set of tokens.
package scanner

import (
	"github.com/kdeconinck/align/internal/pkg/automata/dfa"
	"github.com/kdeconinck/align/internal/pkg/automata/nfa"
)

// ScannerBuilder implements a fluent API for creating a [Scanner].
type ScannerBuilder[S comparable, V any] struct {
	patterns []pattern[S, V]
}

type pattern[S comparable, V any] struct {
	fragment Fragment[S, V]
	value    V
}

func NewScannerBuilder[S comparable, V any]() *ScannerBuilder[S, V] {
	return &ScannerBuilder[S, V]{}
}

func (b *ScannerBuilder[S, V]) Add(fragment Fragment[S, V], value V) *ScannerBuilder[S, V] {
	b.patterns = append(b.patterns, pattern[S, V]{fragment: fragment, value: value})
	return b
}

func (b *ScannerBuilder[S, V]) Build(defaultValue, finalValue V) *Scanner[S, V] {
	n := nfa.New[S, V]()
	start := n.Start()

	for _, p := range b.patterns {
		end := p.fragment.Build(n, start)
		n.AddAcceptingEpsilonTransition(end, p.value)
	}

	return &Scanner[S, V]{
		machine: dfa.FromNfa(n),
		illegal: defaultValue,
		eof:     finalValue,
	}
}
