package adoc

type CodeBlock []Element

func (cb CodeBlock) Text() string {
	txt := ""
	for _, e := range cb {
		txt += e.Text()
	}

	return txt
}

func (p *Parser) tryCodeBlock() bool {
	delimiter := Token{Type: TYPE_DASH, Literal: "----"}
	toks, ok := p.readN(2)
	if !ok {
		return false
	}
	if !toks[0].Equal(delimiter) || toks[1].Type != TYPE_NEWLINE {
		p.unread(2)
		return false
	}
	for {
		tok, ok := p.read()
		if !ok {
			p.unread(len(toks))
			return false
		}
		if tok.Equal(delimiter) {
			break
		}
		toks = append(toks, tok)
	}

	cb := CodeBlock{}
	for _, tok := range toks[2:] {
		cb = append(cb, p.makePlain(tok))
	}
	p.appendElement(cb)

	return true
}
