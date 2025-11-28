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

import "github.com/kdeconinck/align/internal/pkg/automata/dfa"

// Scanner consumes an input sequence and emits tokens using a DFA.
type Scanner[S comparable, V any] struct {
	machine *dfa.Dfa[S, V]
	illegal V
	eof     V

	input []S
	pos   int
}

// Reset assigns new input to the scanner and rewinds the cursor.
func (s *Scanner[S, V]) Reset(input []S) {
	s.input = input
	s.pos = 0
}

// NextToken returns the next token from the input.
// Behavior:
//   - If no input remains, returns the configured EOF token.
//   - It returns the longest matching token; if none matches, it consumes one symbol and returns the Illegal token.
func (s *Scanner[S, V]) NextToken() V {
	if s.pos >= len(s.input) {
		return s.eof
	}

	curr := s.machine.Start()
	lastAcceptIdx := -1
	var lastAcceptVal V
	lastAcceptPos := s.pos

	for idx := s.pos; idx < len(s.input); idx++ {
		next := curr.OutgoingFor(s.input[idx])
		if next == nil {
			break
		}

		curr = next

		if curr.IsAccepting() {
			lastAcceptIdx = curr.AcceptIdx()
			lastAcceptVal = curr.AcceptValue()
			lastAcceptPos = idx + 1
		}
	}

	if lastAcceptIdx != -1 {
		s.pos = lastAcceptPos
		return lastAcceptVal
	}

	// No match found: consume one symbol and report Illegal.
	s.pos++
	return s.illegal
}
