package element

import (
	"ewintr.nl/adoc/token"
)

type List []ListItem

func (l List) Text() string {
	txt := ""
	for _, li := range l {
		txt += li.Text()
	}

	return txt
}

func (l List) Append(els []Element) Element {
	for _, el := range els {
		if val, ok := el.(ListItem); ok {
			l = append(l, val)
		}
	}
	return l
}

func NewListFromTokens(tr ReadUnreader) (ParseResult, bool) {
	toks, ok := tr.Read(2)
	if !ok {
		return ParseResult{}, false
	}
	if !toks[0].Equal(token.TOKEN_ASTERISK) || toks[1].Type != token.TYPE_WHITESPACE {
		tr.Unread(2)
		return ParseResult{}, false
	}

	tr.Unread(2)
	toks = []token.Token{}
	for {
		ntoks, ok := tr.Read(1)
		if !ok {
			tr.Unread(len(toks))
			return ParseResult{}, false
		}
		tok := ntoks[0]
		if tok.Equal(token.TOKEN_DOUBLE_NEWLINE, token.TOKEN_EOF) {
			break
		}
		toks = append(toks, tok)
	}

	return ParseResult{
		Element: List{},
		Inner:   toks,
	}, true
}
