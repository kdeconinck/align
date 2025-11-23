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

// A builder for creating a [Dfa] from a [nfa.Nfa] using the "Subset Construction" algorithm.
type dfaBuilder[S comparable, V any] struct {
	dfa                 *Dfa[S, V]
	workingQueue        *queue.Queue[[]*nfa.State[S, V]]
	subsetKeyToStateMap map[string]*State[S, V]
}

// Returns a [Dfa] that's equivalent to nfa.
func (builder *dfaBuilder[S, V]) buildFromNfa(n *nfa.Nfa[S, V]) *Dfa[S, V] {
	startStates := findPossibleStates(n.Start())

	builder.dfa.start = builder.buildStartState(startStates)

	for builder.workingQueue.Len() > 0 {
		currentSubset, _ := builder.workingQueue.Dequeue()

		sKey := calculateStatesKey(currentSubset)
		from := builder.subsetKeyToStateMap[sKey]

		for sym, nextSubset := range expandStatesPerSymbol(currentSubset) {
			to := builder.ensureState(nextSubset)
			from.transitions[sym] = to
		}
	}

	return builder.dfa
}

// Build a [State] from states.
// The states parameter is added to the builder's working queue for further expansion.
func (builder *dfaBuilder[S, V]) buildStartState(states []*nfa.State[S, V]) *State[S, V] {
	acceptingIdx, acceptingValue := findAcceptanceIdx(states)

	sState := &State[S, V]{
		id:          0, // start is always 0
		transitions: make(map[S]*State[S, V]),
		acceptIdx:   acceptingIdx,
		value:       acceptingValue,
	}

	sKey := calculateStatesKey(states)

	builder.subsetKeyToStateMap[sKey] = sState
	builder.workingQueue.Enqueue(states)

	return sState
}

// Adds states to the [Dfa] that's being constructed by the builder if it hasn't seen by the builder yet.
func (builder *dfaBuilder[S, V]) ensureState(states []*nfa.State[S, V]) *State[S, V] {
	sKey := calculateStatesKey(states)

	if state, ok := builder.subsetKeyToStateMap[sKey]; ok {
		return state
	}

	acceptingIdx, acceptingValue := findAcceptanceIdx(states)

	if acceptingIdx > -1 {
		state := builder.dfa.newAcceptingState(acceptingIdx, acceptingValue)
		builder.subsetKeyToStateMap[sKey] = state
		builder.workingQueue.Enqueue(states)

		return state
	}

	state := builder.dfa.newState()
	builder.subsetKeyToStateMap[sKey] = state
	builder.workingQueue.Enqueue(states)

	return state
}
