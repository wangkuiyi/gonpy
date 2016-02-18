package gonpy

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
