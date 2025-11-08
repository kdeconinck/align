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

// Package pos provides data structures and utilities for tracking positions within an input stream or file.
package pos

import "fmt"

// Position represents a single location in the input data, typically defined by a line and column number.
type Position struct {
	// Line is the actual line number. The first line is 1.
	Line int

	// Column is the actual column number. The first column is 1.
	Column int
}

// New returns a new [Position].
func New() Position {
	return Position{
		Line:   1,
		Column: 1,
	}
}

// Advance increases the Line (if required) and Column values of the position based on r.
func (p *Position) Advance(r rune) {
	if r != '\r' && r != '\n' {
		p.Column += 1
	}

	if r == '\n' {
		p.Column = 1
		p.Line += 1
	}
}

// String returns the human-readable representation of the position.
func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}
