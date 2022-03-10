package adoc

import (
	"time"
)

func (p *Parser) tryHeader() bool {
	toks, ok := p.readN(2)
	if !ok {
		return false
	}
	if !toks[0].Equal(Token{Type: TYPE_EQUALSIGN, Literal: "="}) {
		p.unread(2)
		return false
	}
	if toks[1].Type != TYPE_WHITESPACE {
		p.unread(2)
		return false
	}
	for {
		tok, ok := p.read()
		if !ok {
			p.unread(len(toks))
			return false
		}
		if tok.Equal(TOKEN_EOF, TOKEN_DOUBLE_NEWLINE) {
			break
		}
		toks = append(toks, tok)
	}

	lines := Split(toks[2:], TYPE_NEWLINE)
	p.doc.Title = Literals(lines[0])

	for _, line := range lines[1:] {
		switch {
		case p.tryHeaderDate(line):
			continue
		case p.tryHeaderField(line):
			continue
		default:
			p.doc.Author = Literals(line)
		}
	}

	return true
}

func (p *Parser) tryHeaderField(line []Token) bool {
	if len(line) < 4 {
		return false
	}
	pair := Split(line, TYPE_WHITESPACE)
	if len(pair) != 2 {
		return false
	}
	key, value := pair[0], pair[1]

	if !HasPattern(key, []TokenType{TYPE_COLON, TYPE_WORD, TYPE_COLON}) {
		return false
	}

	p.doc.Attributes[key[1].Literal] = Literals(value)

	return true
}

func (p *Parser) tryHeaderDate(line []Token) bool {
	date, err := time.Parse("2006-01-02", Literals(line))
	if err != nil {
		return false
	}
	p.doc.Date = date

	return true
}
