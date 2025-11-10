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

// Package nfa implements a non-deterministic finite automaton.
package nfa

import "github.com/kdeconinck/align/internal/pkg/collections/mvmap"

// State is a node in an [Nfa].
type State[S comparable, V any] struct {
	id           int
	edge         edge[S, V]                    // Fast path. Used when there's only a single transition.
	transitions  *mvmap.MvMap[S, *State[S, V]] // Slow path: Used when there are multiple transitions.
	eTransitions []*State[S, V]
	acceptIdx    int
	value        V // The accepting value (if any).
}

// An 'edge' is a "single" transition from on [State] to another.
//
// Reasoning:
// When building a linear chain of transitions (which means that each symbol has 1 outgoing transition), it doesn't make
// a lot of sense to allocate a [mvmap.MvMap].
// Instead of using the [mvmap.MvMap], we use this struct for [State]s that only have a single transition.
// This reduces the amount of allocations because we don't have to allocate a [mvmap.MvMap] for each [State].
type edge[S comparable, V any] struct {
	has bool
	sym S
	to  *State[S, V]
}

// Returns a new [edge] representing a transition on symbol to to.
func newEdge[S comparable, V any](sym S, to *State[S, V]) edge[S, V] {
	return edge[S, V]{
		has: true,
		sym: sym,
		to:  to,
	}
}

// Resets the edge on the state.
func (s *State[S, V]) resetEdge() {
	s.edge = edge[S, V]{
		// NOTE: Intentionally left blank.
	}
}

// ID returns the unique, builder-assigned identifier (starting at 0).
func (s *State[S, V]) ID() int {
	return s.id
}

// OutgoingSymbols returns all the symbols that have at least one outgoing transition from this state.
// Note: The order is undefined.
func (s *State[S, V]) OutgoingSymbols() []S {
	if s.transitions == nil {
		if !s.edge.has {
			return nil
		}

		return []S{
			s.edge.sym,
		}
	}

	return s.transitions.Keys()
}

// OutgoingFor returns all the [State]s reachable from the state by consuming symbol or nil if there are no transitions.
func (s *State[S, V]) OutgoingFor(symbol S) []*State[S, V] {
	if s.transitions == nil {
		if s.edge.has && s.edge.sym == symbol {
			return []*State[S, V]{
				s.edge.to,
			}
		}

		return nil
	}

	symbolTransitions := s.transitions.Get(symbol)
	out := make([]*State[S, V], len(symbolTransitions))
	copy(out, symbolTransitions)

	return out
}

// AcceptIdx returns the acceptance index of the state.
func (s *State[S, V]) AcceptIdx() int {
	return s.acceptIdx
}

// IsAccepting reports whether the state is accepting.
func (s *State[S, V]) IsAccepting() bool {
	return s.AcceptIdx() > -1
}

// AcceptValue returns the accepting value if the state is accepting or the zero value of V if the state is NOT
// accepting.
func (s *State[S, V]) AcceptValue() V {
	var defaultValue V

	if s.acceptIdx <= -1 {
		return defaultValue
	}

	return s.value
}

// Add a transition from s on sym to to.
// When there's NO edge on s, the transition is added as an [edge] (fast path) instead of using the [mvmap.MvMap].
func (s *State[S, V]) put(sym S, to *State[S, V]) {
	if !s.edge.has {
		s.edge = newEdge(sym, to)

		return
	}

	if s.edge.has {
		if s.transitions == nil {
			s.transitions = mvmap.New[S, *State[S, V]]()
		}

		s.transitions.Put(s.edge.sym, s.edge.to)
		s.resetEdge()
	}

	s.transitions.Put(sym, to)
}
