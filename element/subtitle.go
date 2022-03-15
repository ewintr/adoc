package element

import "ewintr.nl/adoc/token"

type SubTitle string

func (st SubTitle) Text() string               { return string(st) }
func (st SubTitle) Append(_ []Element) Element { return st }

type SubSubTitle string

func (st SubSubTitle) Text() string               { return string(st) }
func (st SubSubTitle) Append(_ []Element) Element { return st }

func NewSubTitleFromTokens(tr ReadUnreader) (ParseResult, bool) {
	toks, ok := tr.Read(2)
	if !ok {
		return ParseResult{}, false
	}
	if toks[0].Type != token.TYPE_EQUALSIGN || toks[0].Len() < 2 || toks[1].Type != token.TYPE_WHITESPACE {
		tr.Unread(2)
		return ParseResult{}, false
	}

	for {
		ntoks, ok := tr.Read(1)
		if !ok {
			tr.Unread(len(toks))
			return ParseResult{}, false
		}
		tok := ntoks[0]
		if (tok.Type == token.TYPE_NEWLINE && tok.Len() > 1) || tok.Equal(token.TOKEN_EOF) {
			break
		}
		toks = append(toks, tok)
	}

	var title string
	for _, tok := range toks[2:] {
		if tok.Type == token.TYPE_NEWLINE {
			continue
		}
		title += MakePlain(tok).Text()
	}

	var el Element
	switch toks[0].Len() {
	case 2:
		el = SubTitle(title)
	case 3:
		el = SubSubTitle(title)
	default:
		// ignore lower levels for now
		tr.Unread(len(toks))
		return ParseResult{}, false
	}

	return ParseResult{
		Element: el,
		Inner:   []token.Token{},
	}, true
}
