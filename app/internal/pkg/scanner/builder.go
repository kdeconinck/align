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
	"github.com/kdeconinck/align/internal/pkg/pos"
)

// ScannerBuilder is a tool for constructing a [Scanner] by adding various patterns.
// S is the type of the symbols in the input (e.g., byte, rune) and  V is the type of the returned value.
type ScannerBuilder[S comparable, V any] struct {
	patterns []pattern[S, V]
}

// Associates a [Fragment] (a regular expression building block) with the value it should return upon a match.
type pattern[S comparable, V any] struct {
	fragment Fragment[S, V]
	value    V
}

// NewScannerBuilder creates a new, empty [ScannerBuilder].
func NewScannerBuilder[S comparable, V any]() *ScannerBuilder[S, V] {
	return &ScannerBuilder[S, V]{}
}

// Add appends a new pattern to the builder. It takes a [Fragment] (the pattern to match) and the value to return on a
// successful match. It returns the builder itself for method chaining.
func (builder *ScannerBuilder[S, V]) Add(fragment Fragment[S, V], value V) *ScannerBuilder[S, V] {
	builder.patterns = append(builder.patterns, pattern[S, V]{fragment: fragment, value: value})
	return builder
}

// Build finalizes the construction, converting all added patterns into a fully functional and optimized [Scanner].
// The value to return when NO pattern matches is defaultValue.
// The value to return when the input is exhausted on finalValue.
func (builder *ScannerBuilder[S, V]) Build(defaultValue, finalValue V) *Scanner[S, V] {
	machine := nfa.New[S, V]()
	sState := machine.Start()

	for _, pattern := range builder.patterns {
		pEndState := pattern.fragment.Build(machine, sState)

		machine.AddAcceptingEpsilonTransition(pEndState, pattern.value)
	}

	return &Scanner[S, V]{
		machine:    dfa.FromNfa(machine),
		illegal:    defaultValue,
		eof:        finalValue,
		currentPos: pos.New(),
	}
}
