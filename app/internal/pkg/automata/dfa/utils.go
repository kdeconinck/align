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
	"sort"
	"strconv"
	"strings"

	"github.com/kdeconinck/align/internal/pkg/automata/nfa"
	"github.com/kdeconinck/align/internal/pkg/collections/queue"
	"github.com/kdeconinck/align/internal/pkg/collections/set"
)

// Returns all the possible [nfa.State]s, reachable from states, by following zero or more epsilon transitions.
func findPossibleStates[S comparable, V any](states ...*nfa.State[S, V]) []*nfa.State[S, V] {
	workingQueue := queue.New[*nfa.State[S, V]]()
	reachableStates := make([]*nfa.State[S, V], 0, len(states))

	for _, s := range states {
		workingQueue.Enqueue(s)
		reachableStates = append(reachableStates, s)
	}

	for workingQueue.Len() > 0 {
		queuedState, _ := workingQueue.Dequeue()

		for _, nState := range queuedState.Epsilon() {
			workingQueue.Enqueue(nState)
			reachableStates = append(reachableStates, nState)
		}
	}

	return reachableStates
}

// Returns the acceptance index (and value) in states with the lowest value.
func findAcceptanceIdx[S comparable, V any](states []*nfa.State[S, V]) (int, V) {
	var valueV V
	bestIdx := -1

	for _, s := range states {
		idx := s.AcceptIdx()

		if idx == -1 {
			continue
		}

		if bestIdx == -1 || idx < bestIdx {
			bestIdx = idx
			valueV = s.AcceptValue()
		}
	}

	return bestIdx, valueV
}

// Returns a canonical key for states.
// NOTE: The key is calculated by sorting the IDs of states and joining them by ','.
func calculateStatesKey[S comparable, V any](states []*nfa.State[S, V]) string {
	seen := make(map[int]struct{}, len(states))
	stateIDs := make([]int, 0, len(states))

	for _, state := range states {
		id := state.ID()

		seen[id] = struct{}{}
		stateIDs = append(stateIDs, id)
	}

	sort.Ints(stateIDs)

	var b strings.Builder

	for idx, id := range stateIDs {
		if idx > 0 {
			b.WriteByte(',')
		}

		b.WriteString(strconv.Itoa(id))
	}

	return b.String()
}

// Returns all reachable [nfa.State]s (grouped by symbol), reachable from states.
func expandStatesPerSymbol[S comparable, V any](states []*nfa.State[S, V]) map[S][]*nfa.State[S, V] {
	alphabet := set.New[S]()

	for _, state := range states {
		for _, symbol := range state.OutgoingSymbols() {
			alphabet.Add(symbol)
		}
	}

	statesPerSymbol := make(map[S][]*nfa.State[S, V])

	for _, sym := range alphabet.Values() {
		symbolStates := findReachableStatesForSymbol(states, sym)
		epsilonStates := findPossibleStates(symbolStates...)

		if len(epsilonStates) > 0 {
			statesPerSymbol[sym] = epsilonStates
		}
	}

	return statesPerSymbol
}

// Returns all the possible [nfa.State]s (starting from states) that are reachable by following transitions for symbol.
func findReachableStatesForSymbol[S comparable, V any](states []*nfa.State[S, V], symbol S) []*nfa.State[S, V] {
	var reachableStates []*nfa.State[S, V]

	for _, state := range states {
		reachableStates = append(reachableStates, state.OutgoingFor(symbol)...)
	}

	return reachableStates
}
