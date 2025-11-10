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

// QA: Verify and measure the performance of the public API of the "nfa" package.
package nfa_test

import (
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
	"github.com/kdeconinck/align/internal/pkg/automata/nfa"
)

// UT: Create a new [nfa.Nfa].
func TestNew(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[string, int]()
	sState := machine.Start()

	t.Run("When constructing an 'Nfa', it contains NO elements.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := sState.OutgoingSymbols()

		// Assert.
		assert.IsEmptyf(t, got, "\n\n"+
			"UT Name:  When constructing an 'Nfa', it contains NO elements.\n"+
			"\033[32mExpected: 0.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", len(got))
	})

	t.Run("When constructing an 'Nfa', the start state has the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := sState.ID(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When constructing an 'Nfa', the start state has the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When constructing an 'Nfa', the start state is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := sState.IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  When constructing an 'Nfa', the start state is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Build an [nfa.Nfa] and add a single 'accepting' transition.
func TestNfa_Build(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[string, int]()
	sState := machine.Start()
	machine.AddAccepting(sState, "a", 10)

	t.Run("The start 'State' has 1 outgoing symbol ('a').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := sState.OutgoingSymbols(), newSlice("a")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('a').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		states := sState.OutgoingFor("a")

		// Assert.
		got, want := len(states), 1

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State'.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[0].ID(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 accepting 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 accepting 'State'.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got := aStates[0].IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 accepting 'State'.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State' with the correct accepting index.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct accepting index.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[0].AcceptIdx(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct accepting index.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State' with the correct accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct accepting value.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[0].AcceptValue(), 10

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct accepting value.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Build an [nfa.Nfa] and branch on the same symbol into 2 different paths (accepting and NOT accepting).
func TestNfa_BuildWithBracnhingOnTheSameSymbol(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[string, int]()
	sState := machine.Start()
	machine.Add(sState, "a")
	machine.AddAccepting(sState, "a", 10)

	t.Run("The start 'State' has 1 outgoing symbol ('a').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := sState.OutgoingSymbols(), newSlice("a")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('a').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 2 'State's.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		states := sState.OutgoingFor("a")

		// Assert.
		got, want := len(states), 2

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 2 'State's with the correct ID for the 1st.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 2, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's with the correct ID for the 1st.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[0].ID(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's with the correct ID for the 1st.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 2 'State's with the correct ID for the 2nd.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 2, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's with the correct ID for the 2nd.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[1].ID(), 2

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's with the correct ID for the 2nd.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 2 'State's where the 1st one is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 2, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 1st one is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got := aStates[0].IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 1st one is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one is accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 2, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one is accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got := aStates[1].IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one is accepting.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one has the correct accepting index.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 2, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one has the correct accepting index.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[1].AcceptIdx(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one has the correct accepting index.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one has the correct accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 2, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one has the correct accepting value.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[1].AcceptValue(), 10

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 2 'State's where the 2nd one has the correct accepting value.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})
}

// UT: Build an [nfa.Nfa] using a sequence.
func TestNfa_BuildWithSequence(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[string, int]()
	sState := machine.Start()
	aState := machine.Add(sState, "a")
	bState := machine.Add(aState, "b")
	machine.AddAccepting(bState, "c", 10)

	t.Run("The start 'State' has 1 outgoing symbol ('a').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := sState.OutgoingSymbols(), newSlice("a")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('a').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		states := sState.OutgoingFor("a")

		// Assert.
		got, want := len(states), 1

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State'.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[0].ID(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State' that is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got := aStates[0].IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbols ['a', 'b'] leads to exactly 1 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State'.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		states := aStates[0].OutgoingFor("b")

		// Assert.
		got, want := len(states), 1

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State'.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'b'] leads to exactly 1 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		bStates := aStates[0].OutgoingFor("b")

		assert.Truef(t, len(bStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming ['a', 'b'] should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := bStates[0].ID(), 2

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'b'] leads to exactly 1 'State' that is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		bStates := aStates[0].OutgoingFor("b")

		assert.Truef(t, len(bStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming ['a', 'b'] should lead to exactely 1 'State'.\033[0m\n\n")

		got := bStates[0].IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State'.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		bStates := aStates[0].OutgoingFor("b")

		assert.Truef(t, len(bStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State'.\n"+
			"\033[31mFatal error: Consuming ['a', 'b'] should lead to exactely 1 'State'.\033[0m\n\n")

		states := bStates[0].OutgoingFor("c")

		// Assert.
		got, want := len(states), 1

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State'.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		bStates := aStates[0].OutgoingFor("b")

		assert.Truef(t, len(bStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming ['a', 'b'] should lead to exactely 1 'State'.\033[0m\n\n")

		cStates := bStates[0].OutgoingFor("c")

		assert.Truef(t, len(cStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming ['a', 'b', 'c'] should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := cStates[0].ID(), 3

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' that is accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' that is accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		bStates := aStates[0].OutgoingFor("b")

		assert.Truef(t, len(bStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' that is accepting.\n"+
			"\033[31mFatal error: Consuming ['a', 'b'] should lead to exactely 1 'State'.\033[0m\n\n")

		cStates := bStates[0].OutgoingFor("c")

		assert.Truef(t, len(cStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' that is accepting.\n"+
			"\033[31mFatal error: Consuming ['a', 'b', 'c'] should lead to exactely 1 'State'.\033[0m\n\n")

		got := cStates[0].IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to exactly 1 'State' that is accepting.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Build an [nfa.Nfa] using epsilon transitions.
func TestNfa_BuildWithEpsilons(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[string, int]()
	sState := machine.Start()
	e1State := machine.AddAcceptingEpsilonTransition(sState, 10)
	e2State := machine.AddEpsilonTransition(sState)
	machine.AddAccepting(sState, "x", 100)

	t.Run("There are exactely 2 epsilon transitions.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := sState.Epsilon(), newSlice(e1State, e2State)

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  There are exactely 2 epsilon transitions.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("The 1st epsilon transition should be accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		eStates := sState.Epsilon()

		assert.Truef(t, len(eStates) == 2, "\n\n"+
			"UT Name:  The 1st epsilon transition should be accepting.\n"+
			"\033[31mFatal error: The should be 2 epsilon transitions.\033[0m\n\n")

		got := eStates[0].IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  The 1st epsilon transition should be accepting.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("The 2nd epsilon transition should NOT be accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		eStates := sState.Epsilon()

		assert.Truef(t, len(eStates) == 2, "\n\n"+
			"UT Name:  The 2nd epsilon transition should NOT be accepting.\n"+
			"\033[31mFatal error: The should be 2 epsilon transitions.\033[0m\n\n")

		got := eStates[1].IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  The 2nd epsilon transition should NOT be accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Build an [nfa.Nfa] using a loop on an existing [nfa.State].
func TestNfa_BuildWithLoop(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	machine := nfa.New[string, int]()
	sState := machine.Start()
	aState := machine.Add(sState, "a")
	machine.Connect(aState, "a", aState)

	t.Run("The start 'State' has 1 outgoing symbol ('a').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := sState.OutgoingSymbols(), newSlice("a")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('a').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		states := sState.OutgoingFor("a")

		// Assert.
		got, want := len(states), 1

		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State'.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aStates[0].ID(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to exactly 1 'State' that is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		got := aStates[0].IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbols ['a', 'a'] leads to exactly 1 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State'.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		states := aStates[0].OutgoingFor("a")

		got, want := len(states), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State'.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'a'] leads to exactly 1 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		aaStates := aStates[0].OutgoingFor("a")

		assert.Truef(t, len(aaStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming ['a', 'a'] should lead to exactely 1 'State'.\033[0m\n\n")

		got, want := aaStates[0].ID(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'a'] leads to exactly 1 'State' that is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aStates := sState.OutgoingFor("a")

		assert.Truef(t, len(aStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to exactely 1 'State'.\033[0m\n\n")

		aaStates := aStates[0].OutgoingFor("a")

		assert.Truef(t, len(aaStates) == 1, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming ['a', 'a'] should lead to exactely 1 'State'.\033[0m\n\n")

		got := aaStates[0].IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'a'] leads to exactly 1 'State' that is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Verify that copies of elements are returned.
func TestNfa_CopySemantics(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("The 'OutgoingFor' operation does return a copy.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[string, int]()
		sState := machine.Start()
		machine.Add(sState, "a")

		// Act: Mutate the returned value.
		statesFor := sState.OutgoingFor("a")
		statesFor = statesFor[:0]

		// Assert: The mutation (see above) has NO effect on the 'State'.
		got := sState.OutgoingFor("a")

		assert.Equalf(t, len(got), 1, "\n\n"+
			"UT Name:  The 'OutgoingFor' operation does return a copy.\n"+
			"\033[32mExpected (# amount of states for 'a'): %d.\033[0m\n"+
			"\033[31mActual (# amount of states for 'a'):   %d.\033[0m\n\n", 1, len(got))
	})

	t.Run("The 'Epsilon' operation does return <nil> when there are NO epsilon transitions.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[string, int]()
		sState := machine.Start()

		// Act.
		got := sState.Epsilon()

		// Assert.
		assert.Nilf(t, got, "\n\n"+
			"UT Name:  The 'Epsilon' operation does return a copy.\n"+
			"\033[32mExpected: <nil>.\033[0m\n"+
			"\033[31mActual:   NOT <nil>.\033[0m\n\n")
	})

	t.Run("The 'Epsilon' operation does return a copy.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		machine := nfa.New[string, int]()
		sState := machine.Start()
		machine.AddEpsilonTransition(sState)
		machine.AddEpsilonTransition(sState)

		// Act: Mutate the returned value.
		states := sState.Epsilon()
		states = states[:0]

		// Assert: The mutation (see above) has NO effect on the 'State'.
		got := sState.Epsilon()

		assert.Equalf(t, len(got), 2, "\n\n"+
			"UT Name:  The 'Epsilon' operation does return a copy.\n"+
			"\033[32mExpected (# amount of states): %d.\033[0m\n"+
			"\033[31mActual (# amount of states):   %d.\033[0m\n\n", 2, len(got))
	})
}

// UT: Verify that the outgoing symbols for a symbol on a [nfa.State] can be requested.
func TestState_OutgoingFor(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	machine := nfa.New[string, int]()
	sState := machine.Start()

	t.Run("For a 'State' without transitions, the 'OutgoingFor' operation returns <nil>.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := sState.OutgoingFor("a")

		// Assert.
		assert.Nilf(t, got, "\n\n"+
			"UT Name:  For a 'State' without transitions, the 'OutgoingFor' operation returns <nil>.\n"+
			"\033[32mExpected: <nil>.\033[0m\n"+
			"\033[31mActual:   NOT <nil>.\033[0m\n\n")
	})
}

// UT: Verify that the accept value of an [nfa.State] is correct.
func TestState_AcceptValue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	machine := nfa.New[string, int]()
	sState := machine.Start()
	aState := machine.AddAccepting(sState, "a", 10)

	t.Run("For a 'State' that's NOT accepting, the 'AcceptValue' operation returns the default value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := sState.AcceptValue(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  For a 'State' that's NOT accepting, the 'AcceptValue' operation returns the default value.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("For a 'State' that's accepting, the 'AcceptValue' operation returns the accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := aState.AcceptValue(), 10

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  For a 'State' that's accepting, the 'AcceptValue' operation returns the accepting value.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// Utility: Return a slice of T, containing args.
func newSlice[T any](args ...T) []T {
	container := make([]T, len(args))
	copy(container, args)

	return container
}

var benchmarkOutput int // Output of the benchmark(s). Used to avoid compiler optimizations.

// Benchmark(s): Build a chain of transitions in an [nfa.Nfa].
func BenchmarkLinear_1(b *testing.B)         { benchmarkLinear(1, b) }
func BenchmarkLinear_10(b *testing.B)        { benchmarkLinear(10, b) }
func BenchmarkLinear_100(b *testing.B)       { benchmarkLinear(100, b) }
func BenchmarkLinear_1000(b *testing.B)      { benchmarkLinear(1000, b) }
func BenchmarkLinear_1_000_000(b *testing.B) { benchmarkLinear(1_000_000, b) }

// Benchmark(s): Build a chain of epsilon transitions in an [nfa.Nfa].
func BenchmarkLinearEpsilon_1(b *testing.B)         { benchmarkLinearEpsilon(1, b) }
func BenchmarkLinearEpsilon_10(b *testing.B)        { benchmarkLinearEpsilon(10, b) }
func BenchmarkLinearEpsilon_100(b *testing.B)       { benchmarkLinearEpsilon(100, b) }
func BenchmarkLinearEpsilon_1000(b *testing.B)      { benchmarkLinearEpsilon(1000, b) }
func BenchmarkLinearEpsilon_1_000_000(b *testing.B) { benchmarkLinearEpsilon(1_000_000, b) }

// Benchmark(s): Build a "fanout" of transitions in an [nfa.Nfa].
func BenchmarkFanout_1(b *testing.B)         { benchmarkFanout(1, b) }
func BenchmarkFanout_10(b *testing.B)        { benchmarkFanout(10, b) }
func BenchmarkFanout_100(b *testing.B)       { benchmarkFanout(100, b) }
func BenchmarkFanout_1000(b *testing.B)      { benchmarkFanout(1000, b) }
func BenchmarkFanout_1_000_000(b *testing.B) { benchmarkFanout(1_000_000, b) }

// Benchmark: Measure the performance of building a chain of transitions in an [nfa.Nfa].
// Parameters:
// - count: The length of the chain to build.
// - b:     The [testing.B] instance.
func benchmarkLinear(count int, b *testing.B) {
	for b.Loop() {
		machine := nfa.New[byte, int]()
		cState := machine.Start()

		for range count {
			cState = machine.Add(cState, 'a')
		}

		benchmarkOutput = cState.ID()
	}
}

// Benchmark: Measure the performance of building a "fanout" of transitions in an [nfa.Nfa].
// Parameters:
// - fanoutSize: The size of the "fanout" to build.
// - b:          The [testing.B] instance.
func benchmarkFanout(fanoutSize int, b *testing.B) {
	symbols := generateAlphabetSlice(fanoutSize)

	for b.Loop() {
		machine := nfa.New[byte, int]()
		cState := machine.Start()

		for idx := range fanoutSize {
			machine.Add(cState, symbols[idx])
		}

		benchmarkOutput = cState.ID()
	}
}

// Benchmark: Measure the performance of building a chain of epsilon transitions in an [nfa.Nfa].
// Parameters:
// - count: The length of the chain to build.
// - b:     The [testing.B] instance.
func benchmarkLinearEpsilon(count int, b *testing.B) {
	for b.Loop() {
		machine := nfa.New[byte, int]()
		cState := machine.Start()

		for range count {
			cState = machine.AddEpsilonTransition(cState)
		}

		benchmarkOutput = cState.ID()
	}
}

// Returns a slice of n bytes where each byte represents a lowercase letter.
func generateAlphabetSlice(n int) []byte {
	out := make([]byte, n)

	for idx := range n {
		out[idx] = byte((idx % 26) + 'a')
	}

	return out
}
