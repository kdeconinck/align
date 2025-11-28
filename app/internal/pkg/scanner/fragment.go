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

import "github.com/kdeconinck/align/internal/pkg/automata/nfa"

// Fragment is the interface for all components that can build a part an [nfa.Nfa].
type Fragment[S comparable, V any] interface {
	// Build integrates the fragment's logic into the [nfa.Nfa] machine, starting from startState, and returns the final
	// state of the constructed part.
	Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V]
}

// A [Fragment] that matches a fixed sequence of symbols.
type fragLiteral[S comparable, V any] struct {
	symbols []S
}

// A [Fragment] that matches a concatenation of multiple [Fragment]s in order.
type fragSequence[S comparable, V any] struct {
	fragments []Fragment[S, V]
}

// A [Fragment] that matches any one of the provided [Fragment]s.
type fragAnyOf[S comparable, V any] struct {
	fragments []Fragment[S, V]
}

// A [Fragment] that matches a [Fragment] repeated a minimum and (optionally) a maximum number of times.
type fragRepeat[S comparable, V any] struct {
	fragment     Fragment[S, V]
	minOccurence int
	maxOccurence int
	hasMax       bool
}

// Literal creates a [Fragment] that matches the exact, ordered sequence of symbols.
// Panics if no symbols are provided.
func Literal[S comparable, V any](symbols ...S) Fragment[S, V] {
	if len(symbols) == 0 {
		panic("Literal: symbols must have elements")
	}

	return fragLiteral[S, V]{
		symbols: symbols,
	}
}

// Build creates a simple chain of nfa states connected by the literal symbols.
func (frag fragLiteral[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	last := startState

	for _, sym := range frag.symbols {
		last = machine.Add(last, sym)
	}

	return last
}

// Sequence creates a [Fragment] that matches fragments in order.
func Sequence[S comparable, V any](fragments ...Fragment[S, V]) Fragment[S, V] {
	return fragSequence[S, V]{
		fragments: fragments,
	}
}

// Build chains the nfa parts built by the fragments.
func (frag fragSequence[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	last := startState

	for _, sub := range frag.fragments {
		last = sub.Build(machine, last)
	}

	return last
}

// AnyOf creates a [Fragment] that matches any one of fragments.
// Panics if fewer than 2 fragments are provided.
func AnyOf[S comparable, V any](fragments ...Fragment[S, V]) Fragment[S, V] {
	if len(fragments) < 2 {
		panic("AnyOf: at least 2 fragments are required")
	}

	return fragAnyOf[S, V]{
		fragments: fragments,
	}
}

// Build implements the 'OR' construction by creating parallel paths.
// All paths start with an epsilon transition from the start state and end with an epsilon transition to a common end
// state.
func (frag fragAnyOf[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	endState := machine.NewState() // Common 'end' state for each possible fragment.

	for _, sFrag := range frag.fragments {
		branchStart := machine.AddEpsilonTransition(startState)
		branchEnd := sFrag.Build(machine, branchStart)
		machine.ConnectEpsilon(branchEnd, endState)
	}

	return endState
}

// RepeatAtLeast creates a [Fragment] that matches the given fragment at least 'min' times.
// Panics if min is negative.
func RepeatAtLeast[S comparable, V any](min int, fragment Fragment[S, V]) Fragment[S, V] {
	if min < 0 {
		panic("RepeatAtLeast: min cannot be negative")
	}

	return fragRepeat[S, V]{
		fragment:     fragment,
		minOccurence: min,
		maxOccurence: 0,
		hasMax:       false,
	}
}

// RepeatBetween creates a [Fragment] that matches the given fragment between 'min' and 'max' times, inclusive.
// Panics if min is negative or max is less than min.
func RepeatBetween[S comparable, V any](min, max int, fragment Fragment[S, V]) Fragment[S, V] {
	if min < 0 {
		panic("RepeatBetween: min cannot be negative")
	}

	if max < min {
		panic("RepeatBetween: max cannot be less than min")
	}

	return fragRepeat[S, V]{
		fragment:     fragment,
		minOccurence: min,
		maxOccurence: max,
		hasMax:       true,
	}
}

// Build implements the repetition:
// 1. Mandatory `minOccurence` repetitions (a simple sequence).
// 2. Optional `maxOccurence - minOccurence` repetitions, or Kleene star logic if no max.
func (frag fragRepeat[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	currentState := startState

	// Build the sequence that represents the minimal amount of required occurences.
	for idx := 0; idx < frag.minOccurence; idx++ {
		currentState = frag.fragment.Build(machine, currentState)
	}

	endState := machine.NewState()
	machine.ConnectEpsilon(currentState, endState)

	// If there's NO maximal amount of required of occurences, loop back to the beginning.
	if !frag.hasMax {
		machine.ConnectEpsilon(currentState, endState)

		bodyStartState := machine.AddEpsilonTransition(currentState)
		bodyEndState := frag.fragment.Build(machine, bodyStartState)

		machine.ConnectEpsilon(bodyEndState, bodyStartState)
		machine.ConnectEpsilon(bodyEndState, endState)

		return endState
	}

	if frag.maxOccurence-frag.minOccurence == 0 {
		return endState
	}

	for idx := 0; idx < frag.maxOccurence-frag.minOccurence; idx++ {
		optStartState := machine.AddEpsilonTransition(currentState)
		optEndState := frag.fragment.Build(machine, optStartState)

		machine.ConnectEpsilon(optEndState, endState)

		currentState = optEndState
	}

	return endState
}
