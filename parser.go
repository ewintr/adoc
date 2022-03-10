package adoc

import "io"

type Parser struct {
	in  chan Token
	doc *ADoc
	ts  []Token
	r   int
	els []Element
}

func NewParser(reader io.Reader) *Parser {
	lex := NewLexer(reader)
	return NewParserFromChannel(lex.Out())
}

func NewParserFromChannel(toks chan Token) *Parser {
	return &Parser{
		in:  toks,
		doc: NewADoc(),
		ts:  []Token{},
		els: []Element{},
	}
}

func (p *Parser) Parse() *ADoc {
	p.tryHeader()

	for {
		if done := p.scan(); done {
			break
		}
		p.discard()
	}

	p.doc.Content = p.els
	return p.doc
}

func (p *Parser) Elements() []Element {
	return p.els
}

func (p *Parser) scan() bool {
	if _, ok := p.peek(); !ok {
		return true
	}

	type tryFunc func() bool
	for _, tf := range []tryFunc{
		p.tryCodeBlock, p.trySubTitle, p.tryList, p.tryParagraph,
		p.tryStyles, p.tryLink,
	} {
		if tf() {
			return false
		}
	}

	p.doPlain()

	return false
}

func (p *Parser) doPlain() {
	tok, ok := p.read()
	if !ok {
		return
	}
	p.appendElement(p.makePlain(tok))
}

func (p *Parser) makePlain(tok Token) Element {
	switch tok.Type {
	case TYPE_WHITESPACE:
		return WhiteSpace(tok.Literal)
	case TYPE_NEWLINE:
		return WhiteSpace(tok.Literal)
	default:
		return Word(tok.Literal)
	}

}

func (p *Parser) appendElement(b Element) {
	p.els = append(p.els, b)
}

func (p *Parser) consume() bool {
	tok, ok := <-p.in
	if !ok {
		return false
	}
	p.ts = append(p.ts, tok)

	return true
}

func (p *Parser) peek() (Token, bool) {
	if p.r == len(p.ts) {
		if ok := p.consume(); !ok {
			return Token{}, false
		}
	}

	return p.ts[p.r], true
}

func (p *Parser) readN(count int) ([]Token, bool) {
	toks := []Token{}
	for i := 0; i < count; i++ {
		tok, ok := p.read()
		if !ok {
			p.unread(len(toks))
			return []Token{}, false
		}
		toks = append(toks, tok)
	}

	return toks, true
}

func (p *Parser) read() (Token, bool) {
	if p.r == len(p.ts) {
		if ok := p.consume(); !ok {
			return Token{}, false
		}
	}

	tok := p.ts[p.r]
	p.r++

	return tok, true
}

func (p *Parser) unread(count int) {
	for i := count; i > 0; i-- {
		p.r--
	}
}

func (p *Parser) discard() {
	p.ts = p.ts[p.r:]
	p.r = 0
}
