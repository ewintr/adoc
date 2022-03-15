package element

import (
	"ewintr.nl/adoc/token"
)

type Paragraph struct {
	Elements []Element
}

func (p Paragraph) Text() string {
	txt := ""
	for _, el := range p.Elements {
		txt += el.Text()
	}

	return txt
}

func (p Paragraph) Append(els []Element) Element {
	return Paragraph{
		Elements: append(p.Elements, els...),
	}
}

func NewParagraphFromTokens(tr ReadUnreader) (ParseResult, bool) {
	toks := []token.Token{}
	for {
		tok, ok := tr.Read(1)
		if !ok {
			tr.Unread(len(toks))
			return ParseResult{}, false
		}
		if tok[0].Equal(token.TOKEN_EOF, token.TOKEN_DOUBLE_NEWLINE) {
			break
		}
		toks = append(toks, tok[0])
	}

	if len(toks) == 0 {
		return ParseResult{}, false
	}

	return ParseResult{
		Element: Paragraph{Elements: []Element{}},
		Inner:   toks,
	}, true
}
