package adoc

type Paragraph []Element

func (p Paragraph) Text() string {
	txt := ""
	for _, e := range p {
		txt += e.Text()
	}

	return txt
}

func (p *Parser) tryParagraph() bool {
	toks := []Token{}
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

	if len(toks) == 0 {
		return false
	}

	par := NewParserFromChannel(NewTokenStream(toks).Out())
	par.Parse()
	b := Paragraph(par.Elements())

	p.appendElement(b)

	return true
}
