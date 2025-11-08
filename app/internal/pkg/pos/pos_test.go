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

// QA: Verify and measure the performance of the public API of the "pos" package.
package pos_test

import (
	"testing"

	"github.com/kdeconinck/align/internal/pkg/assert"
	"github.com/kdeconinck/align/internal/pkg/pos"
)

// UT: Create a new [pos.Position].
func TestNew(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Act.
	p := pos.New()

	// Assert.
	gotLine, gotCol, wantLine, wantCol := p.Line, p.Column, 1, 1

	assert.Equalf(t, gotLine, wantLine, "\n\n"+
		"UT Name:  When creating a new 'Position', it should start at (1:1).\n"+
		"\033[32mExpected: (%d:%d).\033[0m\n"+
		"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)

	assert.Equalf(t, gotCol, wantCol, "\n\n"+
		"UT Name:  When creating a new 'Position', it should start at (1:1).\n"+
		"\033[32mExpected: (%d:%d).\033[0m\n"+
		"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)
}

// UT: Advance a [pos.Position].
func TestPos_Advance(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When consuming a ' ' rune, the 'Column' should be increased.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		p := pos.New()

		// Act.
		(&p).Advance(' ')

		// Assert.
		gotLine, gotCol, wantLine, wantCol := p.Line, p.Column, 1, 2

		assert.Equalf(t, gotLine, wantLine, "\n\n"+
			"UT Name:  When consuming a ' ' rune, the 'Column' should be increased.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)

		assert.Equalf(t, gotCol, wantCol, "\n\n"+
			"UT Name:  When consuming a ' ' rune, the 'Column' should be increased.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)
	})

	t.Run("When consuming multiple ' ' runes, the 'Column' should be increased multiple times.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		p := pos.New()

		// Act.
		(&p).Advance(' ')
		(&p).Advance(' ')
		(&p).Advance(' ')

		// Assert.
		gotLine, gotCol, wantLine, wantCol := p.Line, p.Column, 1, 4

		assert.Equalf(t, gotLine, wantLine, "\n\n"+
			"UT Name:  When consuming multiple ' ' runes, the 'Column' should be increased multiple times.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)

		assert.Equalf(t, gotCol, wantCol, "\n\n"+
			"UT Name:  When consuming multiple ' ' runes, the 'Column' should be increased multiple times.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)
	})

	t.Run("When consuming a '\n' rune, the 'Line' should be increased.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		p := pos.New()

		// Act.
		(&p).Advance('\n')

		// Assert.
		gotLine, gotCol, wantLine, wantCol := p.Line, p.Column, 2, 1

		assert.Equalf(t, gotLine, wantLine, "\n\n"+
			"UT Name:  When consuming a '\n' rune, the 'Line' should be increased.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)

		assert.Equalf(t, gotCol, wantCol, "\n\n"+
			"UT Name:  When consuming a '\n' rune, the 'Line' should be increased.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)
	})

	t.Run("When consuming multiple '\n' runes, the 'Line' should be increased multiple times.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		p := pos.New()

		// Act.
		(&p).Advance('\n')
		(&p).Advance('\n')
		(&p).Advance('\n')

		// Assert.
		gotLine, gotCol, wantLine, wantCol := p.Line, p.Column, 4, 1

		assert.Equalf(t, gotLine, wantLine, "\n\n"+
			"UT Name:  When consuming multiple '\n' runes, the 'Line' should be increased multiple times.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)

		assert.Equalf(t, gotCol, wantCol, "\n\n"+
			"UT Name:  When consuming multiple '\n' runes, the 'Line' should be increased multiple times.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)
	})

	t.Run("When consuming multiple runes, both the 'Line' and 'Column' should be increased.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		p := pos.New()

		// Act: Consume the runes of the string "Hello\nWorld".
		(&p).Advance('H')
		(&p).Advance('e')
		(&p).Advance('l')
		(&p).Advance('l')
		(&p).Advance('o')
		(&p).Advance('\n')
		(&p).Advance('w')
		(&p).Advance('o')
		(&p).Advance('r')
		(&p).Advance('l')
		(&p).Advance('d')

		// Assert.
		gotLine, gotCol, wantLine, wantCol := p.Line, p.Column, 2, 6

		assert.Equalf(t, gotLine, wantLine, "\n\n"+
			"UT Name:  When consuming multiple runes, both the 'Line' and 'Column' should be increased.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)

		assert.Equalf(t, gotCol, wantCol, "\n\n"+
			"UT Name:  When consuming multiple runes, both the 'Line' and 'Column' should be increased.\n"+
			"\033[32mExpected: (%d:%d).\033[0m\n"+
			"\033[31mActual:   (%d:%d).\033[0m\n\n", wantLine, wantCol, gotLine, gotCol)
	})
}

// UT: Get the human-readable representation of a [pos.Position].
func TestPos_String(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	p := pos.New()

	// Act: Consume the runes of the string "Hello\nWorld".
	(&p).Advance('H')
	(&p).Advance('e')
	(&p).Advance('l')
	(&p).Advance('l')
	(&p).Advance('o')
	(&p).Advance('\n')
	(&p).Advance('w')
	(&p).Advance('o')
	(&p).Advance('r')
	(&p).Advance('l')
	(&p).Advance('d')

	// Assert.
	got, want := p.String(), "2:6"

	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When requesting the human-readable representation of a 'Position', it should be \"Line:Column\".\n"+
		"\033[32mExpected: %v.\033[0m\n"+
		"\033[31mActual:   %v.\033[0m\n\n", want, got)
}
