package header

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strconv"
	"unicode/utf8"
)

type lexer struct {
	*bufio.Reader
	peek rune
}

type token struct {
	val     int
	num     int
	sqstr   string
	boolean bool
}

const (
	// Possible values of token.val:
	eof     = 0
	num     = 57346
	sqstr   = 57347
	boolean = 57348
)

func newLexer(r io.Reader) *lexer {
	return &lexer{
		Reader: bufio.NewReader(r),
		peek:   eof,
	}
}

// Lex returns the next token.
func (x *lexer) lex() token {
	for {
		switch c := x.next(); c {
		case eof:
			return token{val: eof}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.num(c)
		case '{', '}', ':', ',':
			return token{val: int(c)}
		case 'T', 'F':
			return x.bool(c)
		case '\'':
			return x.sqstr()
		case ' ', '\t', '\n', '\r':
		default:
			log.Panicf("unrecognized character %q", c)
		}
	}
}

// Return the next rune for the lexer.
func (x *lexer) next() rune {
	if x.peek != eof {
		r := x.peek
		x.peek = eof // Mark x.peek as empty
		return r
	}

	c, s, e := x.ReadRune()

	if e != nil {
		if e == io.EOF {
			return eof
		} else {
			log.Panic(e)
		}
	} else if c == utf8.RuneError && s == 1 {
		log.Panic("invalid utf8")
	}
	return c
}

// Lex a number.
func (x *lexer) num(c rune) token {
	var b bytes.Buffer
	add(&b, c)

L:
	for {
		switch c = x.next(); c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E':
			add(&b, c)
		default:
			break L
		}
	}

	if c != eof {
		x.peek = c
	}

	numv, e := strconv.Atoi(b.String())
	if e != nil {
		log.Panic(e)
	}

	return token{
		val: num,
		num: numv,
	}
}

func (x *lexer) bool(c rune) token {
	var b bytes.Buffer
	add(&b, c)

	boolv := false

	switch c {
	case 'T':
		for i := 1; i < len("True"); i++ {
			add(&b, x.next())
		}
		if b.String() != "True" {
			log.Panicf("Expecting True, got %v", b.String())
		}
		boolv = true
	case 'F':
		for i := 1; i < len("False"); i++ {
			add(&b, x.next())
		}
		if b.String() != "False" {
			log.Panicf("Expecting False, got %v", b.String())
		}
		boolv = false
	default:
		log.Panicf("Expecting T or F, but got %v", c)
	}

	return token{
		val:     boolean,
		boolean: boolv,
	}
}

func (x *lexer) sqstr() token {
	var b bytes.Buffer

L:
	for {
		switch c := x.next(); c {
		case '\\':
			switch c := x.next(); c {
			case '\\', '\'':
				add(&b, c)
			case 'n':
				add(&b, '\n')
			case 'r':
				add(&b, '\r')
			case 'b':
				add(&b, '\b')
			case 't':
				add(&b, '\t')
			case eof:
				log.Panicf("Unexpected eof in sqstr")
			default:
				add(&b, c)
			}
		case '\'':
			break L
		case eof:
			log.Panicf("Unexpected eof in sqstr")
		default:
			add(&b, c)
		}
	}

	return token{
		val:   sqstr,
		sqstr: b.String(),
	}
}

func add(b *bytes.Buffer, c rune) {
	if _, e := b.WriteRune(c); e != nil {
		log.Panic(e)
	}
}
