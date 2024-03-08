package element

import "code.ewintr.nl/adoc/token"

type Strong []Element

func (st Strong) Text() string {
	var txt string
	for _, e := range st {
		txt += e.Text()
	}

	return txt
}

func (st Strong) Append(els []Element) Element {
	return append(st, els...)
}

type Emphasis []Element

func (em Emphasis) Text() string {
	var txt string
	for _, e := range em {
		txt += e.Text()
	}

	return txt
}

func (em Emphasis) Append(els []Element) Element {
	return append(em, els...)
}

type Code []Element

func (c Code) Text() string {
	var txt string
	for _, e := range c {
		txt += e.Text()
	}

	return txt
}

func (c Code) Append(els []Element) Element {
	return append(c, els...)
}

func NewStyleFromTokens(tr ReadUnreader) (ParseResult, bool) {
	toks, ok := tr.Read(2)
	if !ok {
		return ParseResult{}, false
	}
	markers := []token.Token{
		{Type: token.TYPE_ASTERISK, Literal: "*"},
		{Type: token.TYPE_UNDERSCORE, Literal: "_"},
		{Type: token.TYPE_BACKTICK, Literal: "`"},
	}
	if !toks[0].Equal(markers...) {
		tr.Unread(2)
		return ParseResult{}, false
	}
	if toks[1].Type == token.TYPE_WHITESPACE {
		tr.Unread(2)
		return ParseResult{}, false
	}

	marker := toks[0]
	toks = []token.Token{toks[1]}
	for {
		ntoks, ok := tr.Read(1)
		if !ok {
			tr.Unread(len(toks) + 1)
			return ParseResult{}, false
		}
		tok := ntoks[0]
		if tok.Equal(token.TOKEN_EOF, token.TOKEN_DOUBLE_NEWLINE) {
			tr.Unread(len(toks) + 1)
			return ParseResult{}, false
		}
		if tok.Equal(marker) {
			break
		}
		toks = append(toks, tok)
	}
	if toks[len(toks)-1].Type == token.TYPE_WHITESPACE {
		tr.Unread(len(toks) + 2)
		return ParseResult{}, false
	}

	var s Element
	switch marker.Type {
	case token.TYPE_ASTERISK:
		s = Strong{}
	case token.TYPE_UNDERSCORE:
		s = Emphasis{}
	case token.TYPE_BACKTICK:
		s = Code{}
	}

	return ParseResult{
		Element: s,
		Inner:   toks,
	}, true
}
