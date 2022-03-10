package adoc

type ListItem []Element

func (li ListItem) Text() string {
	txt := ""
	for _, e := range li {
		txt += e.Text()
	}

	return txt
}

type List []ListItem

func (l List) Text() string {
	txt := ""
	for _, li := range l {
		txt += li.Text()
	}

	return txt
}

func (p *Parser) tryList() bool {
	toks, ok := p.readN(2)
	if !ok {
		return false
	}
	if !toks[0].Equal(TOKEN_ASTERISK) || toks[1].Type != TYPE_WHITESPACE {
		p.unread(2)
		return false
	}

	p.unread(2)
	toks = []Token{}
	toksCount := 0
	var items []ListItem
	for {
		tok, ok := p.read()
		if !ok {
			p.unread(toksCount)
			return false
		}
		if tok.Equal(TOKEN_NEWLINE, TOKEN_DOUBLE_NEWLINE, TOKEN_EOF) {
			item, ok := tryListItem(toks)
			if !ok {
				p.unread(len(toks))
				return false
			}
			items = append(items, item)
			toks = []Token{}
			if tok.Equal(TOKEN_DOUBLE_NEWLINE, TOKEN_EOF) {
				break
			}
			continue
		}
		toks = append(toks, tok)
		toksCount++
	}

	p.appendElement(List(items))

	return true
}

func tryListItem(toks []Token) (ListItem, bool) {
	if !toks[0].Equal(TOKEN_ASTERISK) || toks[1].Type != TYPE_WHITESPACE {
		return ListItem{}, false
	}
	stream := NewTokenStream(toks[2:]).Out()
	par := NewParserFromChannel(stream)
	par.Parse()
	b := ListItem(par.Elements())

	return b, true
}
