package header

import "log"

type parser struct {
	*lexer
	peek token
}

func newParser(r runeReader) *parser {
	return &parser{
		lexer: newLexer(r),
		peek:  token{val: eof},
	}
}

func (p *parser) next() token {
	if p.peek.val != eof {
		r := p.peek
		p.peek.val = eof
		return r
	}

	return p.lex()
}

type Shape struct {
	Row, Col int
}

func Parse(r runeReader) map[string]interface{} {
	ret := make(map[string]interface{})
	p := newParser(r)

	expect := func(expected int) token {
		t := p.next()
		if t.val != expected {
			log.Panicf("Expecting %d(%q), got %d(%q)", expected, expected, t.val, t.val)
		}
		return t
	}

	expect(int('{'))

	for {
		key := expect(sqstr).sqstr
		expect(int(':'))

		t := p.lex()
		switch t.val {
		case boolean:
			ret[key] = t.boolean
		case sqstr:
			ret[key] = t.sqstr
		case int('('):
			row := expect(num).num
			expect(int(','))
			if t := p.next(); t.val == int(')') {
				ret[key] = &Shape{Row: row, Col: 1}
			} else {
				p.peek = t
				col := expect(num).num
				expect(int(')'))
				ret[key] = &Shape{Row: row, Col: col}
			}
		}

		expect(int(','))

		if t := p.next(); t.val == int('}') {
			return ret
		} else {
			p.peek = t
		}
	}
}
