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

import (
	"github.com/kdeconinck/align/internal/pkg/automata/nfa"
	"github.com/kdeconinck/align/internal/pkg/collections/queue"
)

// Dfa represents a deterministic finite automaton for symbols of type S.
type Dfa[S comparable, V any] struct {
	start       *State[S, V]
	nextStateID int
}

// FromNfa converts and returns n into an equivalent [Dfa] using the "Subset Construction" algorithm.
func FromNfa[S comparable, V any](n *nfa.Nfa[S, V]) *Dfa[S, V] {
	dfaBuilder := &dfaBuilder[S, V]{
		dfa: &Dfa[S, V]{
			nextStateID: 1,
		},
		workingQueue:        queue.New[[]*nfa.State[S, V]](),
		subsetKeyToStateMap: make(map[string]*State[S, V]),
	}

	return dfaBuilder.buildFromNfa(n)
}

// Start returns the Dfa's start [State].
func (d *Dfa[S, V]) Start() *State[S, V] {
	return d.start
}

// Returns a new [State].
func (d *Dfa[S, V]) newState() *State[S, V] {
	id := d.nextStateID
	d.nextStateID++

	return &State[S, V]{
		id:          id,
		transitions: make(map[S]*State[S, V]),
		acceptIdx:   -1,
	}
}

// Returns a new accepting [State].
func (d *Dfa[S, V]) newAcceptingState(acceptIdx int, value V) *State[S, V] {
	id := d.nextStateID
	d.nextStateID++

	return &State[S, V]{
		id:          id,
		transitions: make(map[S]*State[S, V]),
		acceptIdx:   acceptIdx,
		value:       value,
	}
}
