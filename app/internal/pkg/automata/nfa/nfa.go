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

// Nfa represents a non-deterministic finite automaton for symbols of type S with acceptance metadata of type V.
type Nfa[S comparable, V any] struct {
	start           *State[S, V]
	nextStateID     int
	nextAcceptIndex int
}

// New returns a new [Nfa] for symbols of type S with acceptance metadata of type V.
func New[S comparable, V any]() *Nfa[S, V] {
	return &Nfa[S, V]{
		start: &State[S, V]{
			id:           0,
			eTransitions: nil,
			acceptIdx:    -1,
		},
		nextStateID:     1,
		nextAcceptIndex: 0,
	}
}

// Start returns the start [State] of the nfa.
func (n *Nfa[S, V]) Start() *State[S, V] { return n.start }

// Epsilon returns the reachable [State]s following epsilon transitions from the state.
func (s *State[S, V]) Epsilon() []*State[S, V] {
	if s.eTransitions == nil {
		return nil
	}

	out := make([]*State[S, V], len(s.eTransitions))
	copy(out, s.eTransitions)

	return out
}

// Add adds and returns a new transition starting from s for symbol sym.
// Adding a transition causes a new [State] to be generated.
func (n *Nfa[S, V]) Add(s *State[S, V], sym S) *State[S, V] {
	state := n.newState()
	s.put(sym, state)

	return state
}

// Add adds and returns a new transition starting from s for symbol sym.
// Adding a transition causes a new accepting [State] with metadata value to be generated.
func (n *Nfa[S, V]) AddAccepting(s *State[S, V], sym S, value V) *State[S, V] {
	state := n.newAcceptingState(value)
	s.put(sym, state)

	return state
}

// AddEpsilonTransition adds and returns an epsilon transition starting from s.
// Adding an epsilon transition causes a new [State] to be generated.
func (n *Nfa[S, V]) AddEpsilonTransition(s *State[S, V]) *State[S, V] {
	state := n.newState()
	s.eTransitions = append(s.eTransitions, state)

	return state
}

// AddAcceptingEpsilonTransition adds and returns an epsilon transition starting from s.
// Adding an epsilon transition causes a new accepting [State] with metadata value to be generated.
func (n *Nfa[S, V]) AddAcceptingEpsilonTransition(s *State[S, V], value V) *State[S, V] {
	state := n.newAcceptingState(value)
	s.eTransitions = append(s.eTransitions, state)

	return state
}

// Connect adds a transition on sym from from to to.
func (n *Nfa[S, V]) Connect(from *State[S, V], sym S, to *State[S, V]) {
	from.put(sym, to)
}

// Returns a new [State].
func (n *Nfa[S, V]) newState() *State[S, V] {
	id := n.nextStateID
	n.nextStateID++

	return &State[S, V]{
		id:           id,
		eTransitions: nil,
		acceptIdx:    -1,
	}
}

// Returns a new accepting [State].
func (n *Nfa[S, V]) newAcceptingState(value V) *State[S, V] {
	id := n.nextStateID
	n.nextStateID++

	state := &State[S, V]{
		id:           id,
		eTransitions: nil,
		acceptIdx:    -1,
	}

	n.markAccepting(state, value)

	return state
}

// Marks s as accepting with metadata value.
func (n *Nfa[S, V]) markAccepting(s *State[S, V], value V) {
	idx := n.nextAcceptIndex
	n.nextAcceptIndex++

	s.acceptIdx = idx
	s.value = value
}
