package element

import (
	"code.ewintr.nl/adoc/token"
)

type ListItem []Element

func (li ListItem) Text() string {
	txt := ""
	for _, e := range li {
		txt += e.Text()
	}

	return txt
}

func (li ListItem) Append(els []Element) Element {
	for _, el := range els {
		li = append(li, el)
	}

	return li
}

func NewListItemFromTokens(tr ReadUnreader) (ParseResult, bool) {
	toks, ok := tr.Read(2)
	if !ok {
		return ParseResult{}, false
	}
	if !toks[0].Equal(token.TOKEN_ASTERISK) || toks[1].Type != token.TYPE_WHITESPACE {
		tr.Unread(2)
		return ParseResult{}, false
	}

	toks = []token.Token{}
	for {
		ntoks, ok := tr.Read(1)
		if !ok {
			tr.Unread(len(toks))
			return ParseResult{}, false
		}
		tok := ntoks[0]
		if tok.Equal(token.TOKEN_NEWLINE, token.TOKEN_EOS) {
			break
		}
		toks = append(toks, tok)
	}

	return ParseResult{
		Element: ListItem{},
		Inner:   toks,
	}, true
}
