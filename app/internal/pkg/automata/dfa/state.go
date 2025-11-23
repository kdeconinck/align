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

// Package dfa implements a deterministic finite automaton.
package dfa

// State is a node in a [Dfa].
type State[S comparable, V any] struct {
	id          int
	transitions map[S]*State[S, V]
	acceptIdx   int
	value       V // The accepting value (if any).
}

// ID returns the unique, builder-assigned identifier (starting at 0).
func (s *State[S, V]) ID() int {
	return s.id
}

// OutgoingSymbols returns all the symbols that have an outgoing transition.
func (s *State[S, V]) OutgoingSymbols() []S {
	symbols := make([]S, 0, len(s.transitions))

	for symbol := range s.transitions {
		symbols = append(symbols, symbol)
	}

	return symbols
}

// OutgoingFor returns the [State] reachable by consuming symbol, or nil if none.
func (s *State[S, V]) OutgoingFor(symbol S) *State[S, V] {
	return s.transitions[symbol]
}

// AcceptIdx returns the accepted index.
func (s *State[S, V]) AcceptIdx() int {
	return s.acceptIdx
}

// IsAccepting reports whether the state is accepting.
func (s *State[S, V]) IsAccepting() bool {
	return s.AcceptIdx() > -1
}

// AcceptValue returns the accepting value if any of the default value of V if the state is NOT accepting.
func (s *State[S, V]) AcceptValue() V {
	var defaultValue V

	if s.acceptIdx <= -1 {
		return defaultValue
	}

	return s.value
}
