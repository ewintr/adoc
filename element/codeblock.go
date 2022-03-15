package element

import (
	"ewintr.nl/adoc/token"
)

type CodeBlock []Element

func (cb CodeBlock) Text() string {
	txt := ""
	for _, e := range cb {
		txt += e.Text()
	}

	return txt
}

func (cb CodeBlock) Append(els []Element) Element {
	return CodeBlock{append(cb, els...)}
}

func NewCodeBlockFromTokens(p ReadUnreader) (ParseResult, bool) {
	delimiter := token.Token{Type: token.TYPE_DASH, Literal: "----"}
	toks, ok := p.Read(2)
	if !ok {
		return ParseResult{}, false
	}
	if !toks[0].Equal(delimiter) || toks[1].Type != token.TYPE_NEWLINE {
		p.Unread(2)
		return ParseResult{}, false
	}
	for {
		ntoks, ok := p.Read(1)
		if !ok {
			p.Unread(len(toks))
			return ParseResult{}, false
		}
		tok := ntoks[0]
		if tok.Equal(delimiter) {
			break
		}
		toks = append(toks, tok)
	}

	cb := CodeBlock{}
	for _, tok := range toks[2:] {
		cb = append(cb, MakePlain(tok))
	}

	return ParseResult{
		Element: cb,
		Inner:   []token.Token{},
	}, true
}
