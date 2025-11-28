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

import (
	"github.com/kdeconinck/align/internal/pkg/automata/dfa"
	"github.com/kdeconinck/align/internal/pkg/pos"
)

// Scanner performs a mechine for performing lexical analysis.
type Scanner[S comparable, V any] struct {
	machine    *dfa.Dfa[S, V]
	illegal    V            // The value to return for an unmatchable sequence.
	eof        V            // The value to return when the input is fully consumed.
	currentPos pos.Position // Tracking for the current position in the source.
}

// NextToken reads from rdr from the current position and returns the next token that's matched by a pattern.
func (s *Scanner[S, V]) NextToken(rdr SymbolReader[S]) V {
	currentState := s.machine.Start()
	consumedSymbols := 0    // how many symbols we consumed for this token attempt
	acceptSymbolCount := -1 // number of symbols consumed at last accepting state

	var lastAcceptVal V

	for {
		symbol, err := rdr.ReadSymbol()

		if err != nil {
			break
		}

		consumedSymbols++
		nextState := currentState.OutgoingFor(symbol)

		if nextState == nil {
			break
		}

		currentState = nextState

		if currentState.IsAccepting() {
			acceptSymbolCount = consumedSymbols
			lastAcceptVal = currentState.AcceptValue()
		}
	}

	if consumedSymbols == 0 {
		return s.eof
	}

	if acceptSymbolCount != -1 {
		for idx := acceptSymbolCount; idx < consumedSymbols; idx++ {
			_ = rdr.UnreadSymbol()
		}

		return lastAcceptVal
	}

	for idx := 1; idx < consumedSymbols; idx++ {
		_ = rdr.UnreadSymbol()
	}

	return s.illegal
}
