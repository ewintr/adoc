package adoc

type Strong []Element

func (st Strong) Text() string {
	var txt string
	for _, e := range st {
		txt += e.Text()
	}

	return txt
}

type Emphasis []Element

func (em Emphasis) Text() string {
	var txt string
	for _, e := range em {
		txt += e.Text()
	}

	return txt
}

type Code []Element

func (c Code) Text() string {
	var txt string
	for _, e := range c {
		txt += e.Text()
	}

	return txt
}

func (p *Parser) tryStyles() bool {
	tok, ok := p.readN(2)
	if !ok {
		return false
	}
	markers := []Token{
		{Type: TYPE_ASTERISK, Literal: "*"},
		{Type: TYPE_UNDERSCORE, Literal: "_"},
		{Type: TYPE_BACKTICK, Literal: "`"},
	}
	if !tok[0].Equal(markers...) {
		p.unread(2)
		return false
	}
	if tok[1].Type == TYPE_WHITESPACE {
		p.unread(2)
		return false
	}

	marker := tok[0]
	toks := []Token{tok[1]}
	for {
		tok, ok := p.read()
		if !ok {
			p.unread(len(toks) + 1)
			return false
		}
		if tok.Equal(TOKEN_EOF, TOKEN_DOUBLE_NEWLINE) {
			p.unread(len(toks) + 1)
			return false
		}
		if tok.Equal(marker) {
			break
		}
		toks = append(toks, tok)
	}
	if toks[len(toks)-1].Type == TYPE_WHITESPACE {
		p.unread(len(toks) + 2)
		return false
	}

	par := NewParserFromChannel(NewTokenStream(toks).Out())
	par.Parse()
	var b Element
	switch marker.Type {
	case TYPE_ASTERISK:
		b = Strong(par.Elements())
	case TYPE_UNDERSCORE:
		b = Emphasis(par.Elements())
	case TYPE_BACKTICK:
		b = Code(par.Elements())
	}

	p.appendElement(b)

	return true
}
