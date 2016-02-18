package header

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexNext(t *testing.T) {
	l := newLexer(strings.NewReader(""))
	assert.Equal(t, rune(eof), l.next())

	l = newLexer(strings.NewReader("{},:"))
	assert.Equal(t, '{', l.next())
	assert.Equal(t, '}', l.next())
	assert.Equal(t, ',', l.next())
	assert.Equal(t, ':', l.next())
}

func TestLexNum(t *testing.T) {
	l := newLexer(strings.NewReader("4321 "))
	k := l.num('5')
	assert.Equal(t, num, k.val)
	assert.Equal(t, 54321, k.num)

	assert.Equal(t, ' ', l.next())
	assert.Equal(t, rune(eof), l.next())
}

func TestLexSqstr(t *testing.T) {
	l := newLexer(strings.NewReader("a\\n\\r\\b \\t\\\\'"))
	k := l.sqstr()
	assert.Equal(t, sqstr, k.val)
	assert.Equal(t, "a\n\r\b \t\\", k.sqstr)
}

func TestLexBool(t *testing.T) {
	l := newLexer(strings.NewReader("rue"))
	k := l.bool('T')
	assert.Equal(t, boolean, k.val)
	assert.Equal(t, true, k.boolean)

	l = newLexer(strings.NewReader("alse"))
	k = l.bool('F')
	assert.Equal(t, boolean, k.val)
	assert.Equal(t, false, k.boolean)
}

func TestLexLex(t *testing.T) {
	l := newLexer(strings.NewReader(""))
	assert.Equal(t, eof, l.lex().val)

	l = newLexer(strings.NewReader("\n{  \t'desc':'<f8',}"))
	assert.Equal(t, int('{'), l.lex().val)
	assert.Equal(t, "desc", l.lex().sqstr)
	assert.Equal(t, int(':'), l.lex().val)
	assert.Equal(t, "<f8", l.lex().sqstr)
	assert.Equal(t, int(','), l.lex().val)
	assert.Equal(t, int('}'), l.lex().val)
}
