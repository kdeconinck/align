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

// QA: Verify and measure the performance of the public API of the "scanner" package.
package scanner_test

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
	"github.com/kdeconinck/align/internal/pkg/scanner"
)

// UT: Use a 'Literal' without symbols.
func TestLiteralPanic(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	handler := func() {
		scanner.Literal[rune, string]()
	}

	// Act / assert.
	assert.Panicf(t, handler, "\n\n"+
		"UT Name:  Using a 'Literal' fragment without symbols causes a panic.\n"+
		"\033[32mExpected: The function should 'panic'.\033[0m\n"+
		"\033[31mActual:   The function did NOT 'panic'.\033[0m\n\n")
}

// UT: Use an 'AnyOf' with an invalid value.
func TestAnyOfPanic(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("Using an 'AnyOf' fragment with a negative value causes a panic.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		handler := func() {
			scanner.AnyOf[rune, string]()
		}

		// Act / assert.
		assert.Panicf(t, handler, "\n\n"+
			"UT Name:  Using an 'AnyOf' fragment without a fragment causes a panic.\n"+
			"\033[32mExpected: The function should 'panic'.\033[0m\n"+
			"\033[31mActual:   The function did NOT 'panic'.\033[0m\n\n")
	})

	t.Run("Using an 'AnyOf' fragment with a single fragment causes a panic.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		handler := func() {
			scanner.AnyOf(
				scanner.Literal[rune, string]('p', 'u', 'b', 'l', 'i', 'c'),
			)
		}

		// Act / assert.
		assert.Panicf(t, handler, "\n\n"+
			"UT Name:  Using an 'AnyOf' fragment with a single fragment causes a panic.\n"+
			"\033[32mExpected: The function should 'panic'.\033[0m\n"+
			"\033[31mActual:   The function did NOT 'panic'.\033[0m\n\n")
	})
}

// UT: Use a 'RepeatAtLeast' with an invalid value.
func TestRepeatAtLeastPanic(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	handler := func() {
		scanner.RepeatAtLeast(
			-1,
			scanner.Literal[rune, string](' '),
		)
	}

	// Act / assert.
	assert.Panicf(t, handler, "\n\n"+
		"UT Name:  Using a 'RepeatAtLeast' fragment with a negative value causes a panic.\n"+
		"\033[32mExpected: The function should 'panic'.\033[0m\n"+
		"\033[31mActual:   The function did NOT 'panic'.\033[0m\n\n")
}

// UT: Use a 'RepeatBetween' with an invalid value.
func TestRepeatBetweenPanic(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("Using a 'RepeatBetween' fragment with a negative value causes a panic.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		handler := func() {
			scanner.RepeatBetween(
				-1, 10,
				scanner.Literal[rune, string](' '),
			)
		}

		// Act / assert.
		assert.Panicf(t, handler, "\n\n"+
			"UT Name:  Using a 'RepeatBetween' fragment with a negative value causes a panic.\n"+
			"\033[32mExpected: The function should 'panic'.\033[0m\n"+
			"\033[31mActual:   The function did NOT 'panic'.\033[0m\n\n")
	})

	t.Run("Using a 'RepeatBetween' fragment with a 'max' value that's greater than the 'min' value causes a panic.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		handler := func() {
			scanner.RepeatBetween(
				5, 2,
				scanner.Literal[rune, string](' '),
			)
		}

		// Act / assert.
		assert.Panicf(t, handler, "\n\n"+
			"UT Name:  Using a 'RepeatBetween' fragment with a 'max' value that's greater than the 'min' value causes a panic.\n"+
			"\033[32mExpected: The function should 'panic'.\033[0m\n"+
			"\033[31mActual:   The function did NOT 'panic'.\033[0m\n\n")
	})
}

// UT: Build a [scanner.Scanner] and tokenize a given input.
func TestScanner(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("Scanning when there's NO data produces the final value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := scanner.NewScannerBuilder[rune, string]().
			Build("ILLEGAL", "EOF")

		rdr := strings.NewReader("")
		rRdr := newRuneReader(rdr)

		// Act.
		got, want := readN(s, rRdr, 2), newSlice("EOF", "EOF")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning when there's NO data produces the final value.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Scanning anything that's not recognized produces the default value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		s := scanner.NewScannerBuilder[rune, string]().
			Add(scanner.Literal[rune, string]('a', 'b', 'c'), "abc").
			Add(scanner.Literal[rune, string]('x', 'y', 'z'), "xyz").
			Build("ILLEGAL", "EOF")

		rdr := strings.NewReader("ad")
		rRdr := newRuneReader(rdr)

		// Act.
		got, want := readN(s, rRdr, 3), newSlice("ILLEGAL", "ILLEGAL", "EOF")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning anything that's not recognized produces the default value.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Scanning a 'Literal' fragment produces the value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		frag := scanner.Literal[rune, string]('.', '.')

		s := scanner.NewScannerBuilder[rune, string]().
			Add(frag, "OP").
			Build("ILLEGAL", "EOF")

		rdr := strings.NewReader("..")
		rRdr := newRuneReader(rdr)

		// Act.
		got, want := readN(s, rRdr, 2), newSlice("OP", "EOF")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'Literal' fragment produces the value.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Scanning a 'Sequence' fragment produces the value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		frag := scanner.Sequence(
			scanner.Literal[rune, string]('='),
			scanner.Literal[rune, string]('='),
		)

		s := scanner.NewScannerBuilder[rune, string]().
			Add(frag, "EQUAL").
			Build("ILLEGAL", "EOF")

		rdr := strings.NewReader("==")
		rRdr := newRuneReader(rdr)

		// Act.
		got, want := readN(s, rRdr, 2), newSlice("EQUAL", "EOF")

		// Assert.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'Sequence' fragment produces the value.\n"+
			"\033[32mExpected: %v.\033[0m\n"+
			"\033[31mActual:   %v.\033[0m\n\n", want, got)
	})

	t.Run("Scanning an 'AnyOf' fragment produces the value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		frag := scanner.AnyOf(
			scanner.Literal[rune, string]('p', 'u', 'b', 'l', 'i', 'c'),
			scanner.Literal[rune, string]('P', 'u', 'b', 'l', 'i', 'c'),
			scanner.Literal[rune, string]('P', 'U', 'B', 'L', 'I', 'C'),
		)

		s := scanner.NewScannerBuilder[rune, string]().
			Add(frag, "KW_PUBLIC").
			Build("ILLEGAL", "EOF")

		// Act (reading 'public').
		rdr := strings.NewReader("public")
		rRdr := newRuneReader(rdr)

		got, want := readN(s, rRdr, 2), newSlice("KW_PUBLIC", "EOF")

		// Assert (reading 'public')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning an 'AnyOf' fragment produces the value.\n"+
			"\033[32mExpected (reading 'public'): %v.\033[0m\n"+
			"\033[31mActual (reading 'public'):   %v.\033[0m\n\n", want, got)

		// Act (reading 'Public').
		rdr = strings.NewReader("Public")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 2), newSlice("KW_PUBLIC", "EOF")

		// Assert (reading 'public')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning an 'AnyOf' fragment produces the value.\n"+
			"\033[32mExpected (reading 'Public'): %v.\033[0m\n"+
			"\033[31mActual (reading 'Public'):   %v.\033[0m\n\n", want, got)

		// Act (reading 'PUBLIC').
		rdr = strings.NewReader("PUBLIC")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 2), newSlice("KW_PUBLIC", "EOF")

		// Assert (reading 'PUBLIC')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning an 'AnyOf' fragment produces the value.\n"+
			"\033[32mExpected (reading 'PUBLIC'): %v.\033[0m\n"+
			"\033[31mActual (reading 'PUBLIC'):   %v.\033[0m\n\n", want, got)
	})

	t.Run("Scanning a 'RepeatAtLeast' fragment produces the value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		frag := scanner.RepeatAtLeast(
			2,
			scanner.Literal[rune, string](' '),
		)

		s := scanner.NewScannerBuilder[rune, string]().
			Add(frag, "MULTIPLE_WS").
			Build("ILLEGAL", "EOF")

		// Act (reading ' ').
		rdr := strings.NewReader(" ")
		rRdr := newRuneReader(rdr)

		got, want := readN(s, rRdr, 2), newSlice("ILLEGAL", "EOF")

		// Assert (reading ' ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatAtLeast' fragment produces the value.\n"+
			"\033[32mExpected (reading ' '): %v.\033[0m\n"+
			"\033[31mActual (reading ' '):   %v.\033[0m\n\n", want, got)

		// Act (reading '  ').
		rdr = strings.NewReader("  ")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 2), newSlice("MULTIPLE_WS", "EOF")

		// Assert (reading '  ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatAtLeast' fragment produces the value.\n"+
			"\033[32mExpected (reading '  '): %v.\033[0m\n"+
			"\033[31mActual (reading '  '):   %v.\033[0m\n\n", want, got)

		// Act (reading '   ').
		rdr = strings.NewReader("   ")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 2), newSlice("MULTIPLE_WS", "EOF")

		// Assert (reading '   ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatAtLeast' fragment produces the value.\n"+
			"\033[32mExpected (reading '   '): %v.\033[0m\n"+
			"\033[31mActual (reading '   '):   %v.\033[0m\n\n", want, got)
	})

	t.Run("Scanning a 'RepeatBetween' fragment produces the value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		frag := scanner.RepeatBetween(
			2, 3,
			scanner.Literal[rune, string](' '),
		)

		s := scanner.NewScannerBuilder[rune, string]().
			Add(frag, "MULTIPLE_WS").
			Build("ILLEGAL", "EOF")

		// Act (reading ' ').
		rdr := strings.NewReader(" ")
		rRdr := newRuneReader(rdr)

		got, want := readN(s, rRdr, 2), newSlice("ILLEGAL", "EOF")

		// Assert (reading ' ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatBetween' fragment produces the value.\n"+
			"\033[32mExpected (reading ' '): %v.\033[0m\n"+
			"\033[31mActual (reading ' '):   %v.\033[0m\n\n", want, got)

		// Act (reading '  ').
		rdr = strings.NewReader("  ")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 2), newSlice("MULTIPLE_WS", "EOF")

		// Assert (reading '  ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatBetween' fragment produces the value.\n"+
			"\033[32mExpected (reading '  '): %v.\033[0m\n"+
			"\033[31mActual (reading '  '):   %v.\033[0m\n\n", want, got)

		// Act (reading '   ').
		rdr = strings.NewReader("   ")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 2), newSlice("MULTIPLE_WS", "EOF")

		// Assert (reading '   ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatBetween' fragment produces the value.\n"+
			"\033[32mExpected (reading '   '): %v.\033[0m\n"+
			"\033[31mActual (reading '   '):   %v.\033[0m\n\n", want, got)

		// Act (reading '    ').
		rdr = strings.NewReader("    ")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 3), newSlice("MULTIPLE_WS", "ILLEGAL", "EOF")

		// Assert (reading '    ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatBetween' fragment produces the value.\n"+
			"\033[32mExpected (reading '    '): %v.\033[0m\n"+
			"\033[31mActual (reading '    '):   %v.\033[0m\n\n", want, got)
	})

	t.Run("Scanning a 'RepeatBetween' fragment (with equal 'min' and 'max' values) produces the value.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		frag := scanner.RepeatBetween(
			2, 2,
			scanner.Literal[rune, string](' '),
		)

		s := scanner.NewScannerBuilder[rune, string]().
			Add(frag, "MULTIPLE_WS").
			Build("ILLEGAL", "EOF")

		// Act (reading ' ').
		rdr := strings.NewReader(" ")
		rRdr := newRuneReader(rdr)

		got, want := readN(s, rRdr, 2), newSlice("ILLEGAL", "EOF")

		// Assert (reading ' ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatBetween' fragment (with equal 'min' and 'max' values) produces the value.\n"+
			"\033[32mExpected (reading ' '): %v.\033[0m\n"+
			"\033[31mActual (reading ' '):   %v.\033[0m\n\n", want, got)

		// Act (reading '  ').
		rdr = strings.NewReader("  ")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 2), newSlice("MULTIPLE_WS", "EOF")

		// Assert (reading '  ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatBetween' fragment (with equal 'min' and 'max' values) produces the value.\n"+
			"\033[32mExpected (reading '  '): %v.\033[0m\n"+
			"\033[31mActual (reading '  '):   %v.\033[0m\n\n", want, got)

		// Act (reading '   ').
		rdr = strings.NewReader("   ")
		rRdr = newRuneReader(rdr)

		got, want = readN(s, rRdr, 3), newSlice("MULTIPLE_WS", "ILLEGAL", "EOF")

		// Assert (reading '   ')
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  Scanning a 'RepeatBetween' fragment (with equal 'min' and 'max' values) produces the value.\n"+
			"\033[32mExpected (reading '   '): %v.\033[0m\n"+
			"\033[31mActual (reading '   '):   %v.\033[0m\n\n", want, got)
	})
}

// Read n amount of tokens from scanner.
func readN[S comparable, V any](scanner *scanner.Scanner[S, V], rdr scanner.SymbolReader[S], n int) []V {
	tokens := make([]V, 0, n)

	for idx := 0; idx < n; idx += 1 {
		tokens = append(tokens, scanner.NextToken(rdr))
	}

	return tokens
}

// Utility: Return a slice of T, containing args.
func newSlice[T any](args ...T) []T {
	container := make([]T, len(args))
	copy(container, args)

	return container
}

// A [scanner.SymbolReader] implementation for rune input, backed by a [bufio.Reader].
type runeReader struct {
	rdr *bufio.Reader
}

// Returns a new that [scanner.SymbolReader] that reads runes from rdr.
func newRuneReader(rdr io.Reader) *runeReader {
	return &runeReader{rdr: bufio.NewReader(rdr)}
}

// ReadSymbol reads the next rune.
func (rRdr *runeReader) ReadSymbol() (rune, error) {
	r, _, err := rRdr.rdr.ReadRune()

	return r, err
}

// UnreadSymbol unreads the last rune read.
func (rRdr *runeReader) UnreadSymbol() error {
	return rRdr.rdr.UnreadRune()
}
