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

// Fragment builds part of an NFA starting from startState and returns the end state.
type Fragment[S comparable, V any] interface {
	// Build returns a new [nfa.State] starting from startState on machine.
	Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V]
}

// A fragment that's a set of symbols.
type fragLiteral[S comparable, V any] struct {
	symbols []S
}

// A fragment that's a sequence of other fragments.
type fragSequence[S comparable, V any] struct {
	fragments []Fragment[S, V]
}

// A fragment that matches any of a set of other fragments.
type fragAnyOf[S comparable, V any] struct {
	fragments []Fragment[S, V]
}

// A fragment that matches a repetition of another fragment.
type fragRepeat[S comparable, V any] struct {
	fragment     Fragment[S, V]
	minOccurence int  // The minimal amount of required occurences.
	maxOccurence int  // The maximal amount of occurences. Note: Only valid is 'hasMax' is set to true.
	hasMax       bool // Flag indicating if the fragment can be repeated a maximum amount of times.
}

// Literal builds a chain for the given symbols.
func Literal[S comparable, V any](symbols ...S) Fragment[S, V] {
	return fragLiteral[S, V]{
		symbols: symbols,
	}
}

// Build a literal fragment starting from startState on machine.
func (frag fragLiteral[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	if len(frag.symbols) == 0 {
		return startState
	}

	last := startState

	for _, sym := range frag.symbols {
		last = machine.Add(last, sym)
	}

	return last
}

// Sequence builds a fragment that is the concatenation of fragments.
func Sequence[S comparable, V any](fragments ...Fragment[S, V]) Fragment[S, V] {
	return fragSequence[S, V]{
		fragments: fragments,
	}
}

// Build a sequence of fragments starting from startState on machine.
func (frag fragSequence[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	last := startState

	for _, sub := range frag.fragments {
		last = sub.Build(machine, last)
	}

	return last
}

// AnyOf builds a fragment that is the alternation of fragments.
func AnyOf[S comparable, V any](fragments ...Fragment[S, V]) Fragment[S, V] {
	return fragAnyOf[S, V]{
		fragments: fragments,
	}
}

// Build a sequence of fragments starting from startState on machine.
func (frag fragAnyOf[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	switch len(frag.fragments) {
	case 0:
		return startState
	case 1:
		return frag.fragments[0].Build(machine, startState)
	}

	endState := machine.NewState() // Common 'end' state for each possible fragment.

	for _, sFrag := range frag.fragments {
		branchStart := machine.AddEpsilonTransition(startState)
		branchEnd := sFrag.Build(machine, branchStart)
		machine.ConnectEpsilon(branchEnd, endState)
	}

	return endState
}

// RepeatAtLeast builds a fragment that matches fragment repeated min .. ∞ times.
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

// RepeatBetween builds a fragment that matches fragment repeated min .. max times.
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

// Build a sequence of fragments starting from startState on machine.
func (frag fragRepeat[S, V]) Build(machine *nfa.Nfa[S, V], startState *nfa.State[S, V]) *nfa.State[S, V] {
	current := startState

	for idx := 0; idx < frag.minOccurence; idx++ {
		current = frag.fragment.Build(machine, current)
	}

	if !frag.hasMax {
		endState := machine.NewState()

		// Stop right after the min repetitions.
		machine.ConnectEpsilon(current, endState)

		// One-or-more extra repetitions.
		bodyStart := machine.AddEpsilonTransition(current)
		bodyEnd := frag.fragment.Build(machine, bodyStart)

		// Loop back.
		machine.ConnectEpsilon(bodyEnd, bodyStart)

		// Or exit the loop.
		machine.ConnectEpsilon(bodyEnd, endState)

		return endState
	}

	// Create a common end state.
	end := machine.NewState()

	// We can always stop right after the required min part.
	machine.ConnectEpsilon(current, end)

	// If max == min, we're done: exactly-min repetitions.
	if frag.maxOccurence-frag.minOccurence == 0 {
		return end
	}

	// General bounded case: add up to 'remaining' optional repetitions.
	//
	// Conceptually:
	//
	//   s_m = end of min chain (current)
	//   s_m --ε--> end
	//   s_m --ε--> optStart0 --[inner]--> optEnd0 --ε--> end
	//   optEnd0 --ε--> optStart1 --[inner]--> optEnd1 --ε--> end
	//   ...
	//
	cursor := current

	for i := 0; i < frag.maxOccurence-frag.minOccurence; i++ {
		// From cursor:
		// - ε to end  (stop here)
		// - ε to optStart, then one more 'inner'
		optStart := machine.AddEpsilonTransition(cursor)
		optEnd := frag.fragment.Build(machine, optStart)

		// After this optional repetition, we can stop:
		machine.ConnectEpsilon(optEnd, end)

		// Or do another optional repetition in the next loop:
		cursor = optEnd
	}

	return end
}
