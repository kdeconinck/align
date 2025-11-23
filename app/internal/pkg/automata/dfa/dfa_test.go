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

// QA: Verify and measure the performance of the public API of the "dfa" package.
package dfa_test

import (
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
	"github.com/kdeconinck/align/internal/pkg/automata/dfa"
	"github.com/kdeconinck/align/internal/pkg/automata/nfa"
)

// UT: Create a new [dfa.Dfa] from an empty [nfa.Nfa].
func TestDfa_FromNfa_New(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[string, int]()
	dMachine := dfa.FromNfa(nMachine)
	dSState := dMachine.Start()

	t.Run("When converting an 'Nfa' with NO transitions, the start 'State' has NO outgoing symbols.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := dSState.OutgoingSymbols()

		// Assert.
		assert.IsEmptyf(t, got, "\n\n"+
			"UT Name:  When converting an 'Nfa' with NO transitions, the start 'State' has NO outgoing symbols.\n"+
			"\033[32mExpected (# outgoing symbols): 0.\033[0m\n"+
			"\033[31mActual (# outgoing symbols):   %d.\033[0m\n\n", len(got))
	})

	t.Run("When converting an 'Nfa' with NO transitions, the start 'State' has the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.ID(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When converting an 'Nfa' with NO transitions, the start 'State' has the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("When converting an 'Nfa' with NO transitions, the start 'State' is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := dSState.IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  When converting an 'Nfa' with NO transitions, the start 'State' is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Build an [nfa.Nfa] using a sequence and convert it to a [dfa.Dfa].
func TestDfa_FromNfa_BuildWithSequence(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[string, int]()
	sState := nMachine.Start()
	aState := nMachine.Add(sState, "a")
	bState := nMachine.Add(aState, "b")
	nMachine.AddAccepting(bState, "c", 10)

	dMachine := dfa.FromNfa(nMachine)
	dSState := dMachine.Start()

	t.Run("The start 'State' has 1 outgoing symbol ('a').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.OutgoingSymbols(), newSlice("a")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('a').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to a non-nil 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := dSState.OutgoingFor("a")

		// Assert.
		assert.NotNilf(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a non-nil 'State'.\n"+
			"\033[32mExpected: NOT <nil>.\033[0m\n"+
			"\033[31mActual:   <nil>.\033[0m\n\n")
	})

	t.Run("Consuming the symbol 'a' leads to a 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got, want := aDState.ID(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to a 'State' that is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got := aDState.IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbols ['a', 'b'] leads to a 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State'.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		bDState := aDState.OutgoingFor("b")

		// Assert.
		assert.NotNilf(t, bDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State'.\n"+
			"\033[32mExpected: NOT <nil>.\033[0m\n"+
			"\033[31mActual:   <nil>.\033[0m\n\n")
	})

	t.Run("Consuming the symbols ['a', 'b'] leads to a 'State' with the correct ID.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		bDState := aDState.OutgoingFor("b")

		assert.NotNilf(t, bDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State' with the correct ID.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got, want := bDState.ID(), 2

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State' with the correct ID.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'b'] leads to a 'State' that is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		bDState := aDState.OutgoingFor("b")

		assert.NotNilf(t, bDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got := bDState.IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b'] leads to a 'State' that is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbols ['a', 'b', 'c'] leads to a 'State' that is accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbols Consuming the symbols ['a', 'b', 'c'] leads to a 'State' that is accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		bDState := aDState.OutgoingFor("b")

		assert.NotNilf(t, bDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' that is accepting.\n"+
			"\033[31mFatal error: Consuming 'b' should lead to a 'State'.\033[0m\n\n")

		cDState := bDState.OutgoingFor("c")

		assert.NotNilf(t, cDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' that is accepting.\n"+
			"\033[31mFatal error: Consuming 'c' should lead to a 'State'.\033[0m\n\n")

		got := cDState.IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' that is accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting index.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting index.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		bDState := aDState.OutgoingFor("b")

		assert.NotNilf(t, bDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting index.\n"+
			"\033[31mFatal error: Consuming 'b' should lead to a 'State'.\033[0m\n\n")

		cDState := bDState.OutgoingFor("c")

		assert.NotNilf(t, cDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting index.\n"+
			"\033[31mFatal error: Consuming 'c' should lead to a 'State'.\033[0m\n\n")

		got, want := cDState.AcceptIdx(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting index.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting value.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		bDState := aDState.OutgoingFor("b")

		assert.NotNilf(t, bDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting value.\n"+
			"\033[31mFatal error: Consuming 'b' should lead to a 'State'.\033[0m\n\n")

		cDState := bDState.OutgoingFor("c")

		assert.NotNilf(t, cDState, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting value.\n"+
			"\033[31mFatal error: Consuming 'c' should lead to a 'State'.\033[0m\n\n")

		got, want := cDState.AcceptValue(), 10

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbols ['a', 'b', 'c'] leads to a 'State' with the correct accepting value.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Build an [nfa.Nfa] and branch on the same symbol into 2 different paths (accepting and NOT accepting) convert it
// to a [dfa.Dfa].
func TestDfa_FromNfa_BuildWithBranchingOnTheSameSymbol(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[string, int]()
	sState := nMachine.Start()
	nMachine.Add(sState, "a")
	nMachine.AddAccepting(sState, "a", 10)

	dMachine := dfa.FromNfa(nMachine)
	dSState := dMachine.Start()

	t.Run("The start 'State' has 1 outgoing symbol ('a').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.OutgoingSymbols(), newSlice("a")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('a').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to a non-nil 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := dSState.OutgoingFor("a")

		// Assert.
		assert.NotNilf(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a non-nil 'State'.\n"+
			"\033[32mExpected: NOT <nil>.\033[0m\n"+
			"\033[31mActual:   <nil>.\033[0m\n\n")
	})

	t.Run("Consuming the symbol 'a' leads to a 'State' that is accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that is accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got := aDState.IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that is accepting.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbol 'a' leads to a 'State' with the correct accepting index.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' with the correct accepting index.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got, want := aDState.AcceptIdx(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' with the correct accepting index.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to a 'State' with the correct accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' with the correct accepting value.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got, want := aDState.AcceptValue(), 10

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' with the correct accepting value.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Build an [nfa.Nfa] using epsilon transitions and convert it to a [dfa.Dfa].
func TestDfa_FromNfa_BuildWithEpsilons(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[string, int]()
	sState := nMachine.Start()
	nMachine.AddAcceptingEpsilonTransition(sState, 10)
	eState := nMachine.AddEpsilonTransition(sState)
	nMachine.AddAccepting(eState, "x", 20)

	dMachine := dfa.FromNfa(nMachine)
	dSState := dMachine.Start()

	t.Run("The start 'State' is accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got := dSState.IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  The start 'State' is accepting.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("The start 'State' has the correct accepting index.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.AcceptIdx(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has the correct accepting index.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("The start 'State' has the correct accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.AcceptValue(), 10

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has the correct accepting index.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("The start 'State' has 1 outgoing symbol ('x').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.OutgoingSymbols(), newSlice("x")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('x').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'x' leads to an accepting 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		xDState := dSState.OutgoingFor("x")

		assert.NotNilf(t, xDState, "\n\n"+
			"UT Name:  Consuming the symbol 'x' leads to an accepting 'State'.\n"+
			"\033[31mFatal error: Consuming 'x' should lead to a 'State'.\033[0m\n\n")

		got := xDState.IsAccepting()

		// Assert.
		assert.Truef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'x' leads to an accepting 'State'.\n"+
			"\033[32mExpected: true.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})

	t.Run("Consuming the symbol 'x' leads to a 'State' with the correct accepting index.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		xDState := dSState.OutgoingFor("x")

		assert.NotNilf(t, xDState, "\n\n"+
			"UT Name:  Consuming the symbol 'x' leads to a 'State' with the correct accepting index.\n"+
			"\033[31mFatal error: Consuming 'x' should lead to a 'State'.\033[0m\n\n")

		got, want := xDState.AcceptIdx(), 1

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'x' leads to a 'State' with the correct accepting index.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'x' leads to a 'State' with the correct accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		xDState := dSState.OutgoingFor("x")

		assert.NotNilf(t, xDState, "\n\n"+
			"UT Name:  Consuming the symbol 'x' leads to a 'State' with the correct accepting value.\n"+
			"\033[31mFatal error: Consuming 'x' should lead to a 'State'.\033[0m\n\n")

		got, want := xDState.AcceptValue(), 20

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  Consuming the symbol 'x' leads to a 'State' with the correct accepting value.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})
}

// UT: Build an [nfa.Nfa] using a loop on an existing [nfa.State] and convert it to a [dfa.Dfa].
func TestDfa_FromNfa_BuildWithLoop(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[string, int]()
	sState := nMachine.Start()
	aState := nMachine.Add(sState, "a")
	nMachine.Connect(aState, "a", aState)

	dMachine := dfa.FromNfa(nMachine)
	dSState := dMachine.Start()

	t.Run("The start 'State' has 1 outgoing symbol ('a').", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.OutgoingSymbols(), newSlice("a")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The start 'State' has 1 outgoing symbol ('a').\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Consuming the symbol 'a' leads to a non-nil 'State'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		// Assert.
		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a non-nil 'State'.\n"+
			"\033[32mExpected: NOT <nil>.\033[0m\n"+
			"\033[31mActual:   <nil>.\033[0m\n\n")
	})

	t.Run("Consuming the symbol 'a' leads to a 'State' that loops on itself.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that loops on itself.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got := aDState.OutgoingFor("a")

		// Assert.
		assert.Truef(t, got == aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that loops on itself.\n"+
			"\033[32mExpected: loop to itself.\033[0m\n"+
			"\033[31mActual:   loop to a different 'State'.\033[0m\n\n")
	})

	t.Run("Consuming the symbol 'a' leads to a 'State' that is NOT accepting.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that is NOT accepting.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got := aDState.IsAccepting()

		// Assert.
		assert.Falsef(t, got, "\n\n"+
			"UT Name:  Consuming the symbol 'a' leads to a 'State' that is NOT accepting.\n"+
			"\033[32mExpected: false.\033[0m\n"+
			"\033[31mActual:   %t.\033[0m\n\n", got)
	})
}

// UT: Verify that copies of elements are returned.
func TestDfa_CopySemantics(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("The 'OutgoingSymbols' operation returns a copy.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		nMachine := nfa.New[string, int]()
		sState := nMachine.Start()
		nMachine.Add(sState, "a")

		dMachine := dfa.FromNfa(nMachine)
		dSState := dMachine.Start()

		// Act: Mutate the returned value.
		symbols := dSState.OutgoingSymbols()
		symbols[0] = "x"

		// Assert: The mutation (see above) has NO effect on the 'State'.
		got, want := dSState.OutgoingSymbols(), newSlice("a")

		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  The 'OutgoingSymbols' operation returns a copy.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})
}

// UT: Verify that the accept value of a [dfa.State] is correct.
func TestDfaState_AcceptValue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	nMachine := nfa.New[string, int]()
	sState := nMachine.Start()
	nMachine.AddAccepting(sState, "a", 10)

	dMachine := dfa.FromNfa(nMachine)
	dSState := dMachine.Start()

	t.Run("For a 'State' that's NOT accepting, the 'AcceptValue' operation returns the default value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		got, want := dSState.AcceptValue(), 0

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  For a 'State' that's NOT accepting, the 'AcceptValue' operation returns the default value.\n"+
			"\033[32mExpected: %d.\033[0m\n"+
			"\033[31mActual:   %d.\033[0m\n\n", want, got)
	})

	t.Run("For a 'State' that's accepting, the 'AcceptValue' operation returns the accepting value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Act.
		aDState := dSState.OutgoingFor("a")

		assert.NotNilf(t, aDState, "\n\n"+
			"UT Name:  For a 'State' that's accepting, the 'AcceptValue' operation returns the accepting value.\n"+
			"\033[31mFatal error: Consuming 'a' should lead to a 'State'.\033[0m\n\n")

		got, want := aDState.AcceptValue(), 10

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

// Benchmark(s): Convert a linear [nfa.Nfa] to a [dfa.Dfa].
func BenchmarkFromNfaLinear_1(b *testing.B)         { benchmarkFromNfaLinear(1, b) }
func BenchmarkFromNfaLinear_10(b *testing.B)        { benchmarkFromNfaLinear(10, b) }
func BenchmarkFromNfaLinear_100(b *testing.B)       { benchmarkFromNfaLinear(100, b) }
func BenchmarkFromNfaLinear_1000(b *testing.B)      { benchmarkFromNfaLinear(1000, b) }
func BenchmarkFromNfaLinear_1_000_000(b *testing.B) { benchmarkFromNfaLinear(1_000_000, b) }

// Benchmark(s): Convert a chain of epsilon transitions in an [nfa.Nfa] to a [dfa.Dfa].
func BenchmarkFromNfaLinearEpsilon_1(b *testing.B)    { benchmarkFromNfaLinearEpsilon(1, b) }
func BenchmarkFromNfaLinearEpsilon_10(b *testing.B)   { benchmarkFromNfaLinearEpsilon(10, b) }
func BenchmarkFromNfaLinearEpsilon_100(b *testing.B)  { benchmarkFromNfaLinearEpsilon(100, b) }
func BenchmarkFromNfaLinearEpsilon_1000(b *testing.B) { benchmarkFromNfaLinearEpsilon(1000, b) }
func BenchmarkFromNfaLinearEpsilon_1_000_000(b *testing.B) {
	benchmarkFromNfaLinearEpsilon(1_000_000, b)
}

// Benchmark(s): Convert a "fanout" [nfa.Nfa] to a [dfa.Dfa].
func BenchmarkFromNfaFanout_1(b *testing.B)         { benchmarkFromNfaFanout(1, b) }
func BenchmarkFromNfaFanout_10(b *testing.B)        { benchmarkFromNfaFanout(10, b) }
func BenchmarkFromNfaFanout_100(b *testing.B)       { benchmarkFromNfaFanout(100, b) }
func BenchmarkFromNfaFanout_1000(b *testing.B)      { benchmarkFromNfaFanout(1000, b) }
func BenchmarkFromNfaFanout_1_000_000(b *testing.B) { benchmarkFromNfaFanout(1_000_000, b) }

// Benchmark: Measure the performance of converting a linear [nfa.Nfa] to a [dfa.Dfa].
// Parameters:
// - count: The length of the chain to build.
// - b:     The [testing.B] instance.
func benchmarkFromNfaLinear(count int, b *testing.B) {
	for b.Loop() {
		nMachine := nfa.New[byte, int]()
		cState := nMachine.Start()

		for range count {
			cState = nMachine.Add(cState, 'a')
		}

		dMachine := dfa.FromNfa(nMachine)
		benchmarkOutput = dMachine.Start().ID()
	}
}

// Benchmark: Measure the performance of converting a "fanout" [nfa.Nfa] to a [dfa.Dfa].
// Parameters:
// - fanoutSize: The size of the "fanout" to build.
// - b:          The [testing.B] instance.
func benchmarkFromNfaFanout(fanoutSize int, b *testing.B) {
	symbols := generateAlphabetSlice(fanoutSize)

	for b.Loop() {
		nMachine := nfa.New[byte, int]()
		cState := nMachine.Start()

		for idx := range fanoutSize {
			nMachine.Add(cState, symbols[idx])
		}

		dMachine := dfa.FromNfa(nMachine)
		benchmarkOutput = dMachine.Start().ID()
	}
}

// Benchmark: Measure the performance of converting a chain of epsilon transitions in an [nfa.Nfa] to a [dfa.Dfa].
// Parameters:
// - count: The length of the epsilon chain to build.
// - b:     The [testing.B] instance.
func benchmarkFromNfaLinearEpsilon(count int, b *testing.B) {
	for b.Loop() {
		nMachine := nfa.New[byte, int]()
		cState := nMachine.Start()

		for range count {
			cState = nMachine.AddEpsilonTransition(cState)
		}

		dMachine := dfa.FromNfa(nMachine)
		benchmarkOutput = dMachine.Start().ID()
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
