package adoc

type SubTitle string

func (st SubTitle) Text() string { return string(st) }

type SubSubTitle string

func (st SubSubTitle) Text() string { return string(st) }

func (p *Parser) trySubTitle() bool {
	toks, ok := p.readN(2)
	if !ok {
		return false
	}
	if toks[0].Type != TYPE_EQUALSIGN || toks[0].Len() < 2 || toks[1].Type != TYPE_WHITESPACE {
		p.unread(2)
		return false
	}

	for {
		tok, ok := p.read()
		if !ok {
			p.unread(len(toks))
			return false
		}
		if (tok.Type == TYPE_NEWLINE && tok.Len() > 1) || tok.Equal(TOKEN_EOF) {
			break
		}
		toks = append(toks, tok)
	}

	var title string
	for _, tok := range toks[2:] {
		if tok.Type == TYPE_NEWLINE {
			continue
		}
		title += p.makePlain(tok).Text()
	}
	switch toks[0].Len() {
	case 2:
		p.appendElement(SubTitle(title))
	case 3:
		p.appendElement(SubSubTitle(title))
	default:
		// ignore lower levels for now
		p.unread(len(toks))
		return false
	}

	return true
}
