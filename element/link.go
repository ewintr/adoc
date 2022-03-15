package element

import (
	"ewintr.nl/adoc/token"
)

type Link struct {
	URL   string
	Title string
}

func (l Link) Text() string               { return l.Title }
func (l Link) Append(_ []Element) Element { return l }

func NewLinkFromTokens(tr ReadUnreader) (ParseResult, bool) {
	tok, ok := tr.Read(1)
	if !ok {
		return ParseResult{}, false
	}
	tr.Unread(1)
	if tok[0].Type != token.TYPE_WORD {
		return ParseResult{}, false
	}

	url, title := []token.Token{}, []token.Token{}
	lb := false
	count := 0
	for {
		ntoks, ok := tr.Read(1)
		if !ok {
			tr.Unread(count)
			return ParseResult{}, false
		}
		count++
		tok := ntoks[0]
		if tok.Equal(token.TOKEN_EOF, token.TOKEN_EOS, token.TOKEN_NEWLINE, token.TOKEN_DOUBLE_NEWLINE) {
			tr.Unread(count)
			return ParseResult{}, false
		}
		if !lb && tok.Type == token.TYPE_WHITESPACE {
			tr.Unread(count)
			return ParseResult{}, false
		}
		if tok.Equal(token.Token{Type: token.TYPE_BRACKETOPEN, Literal: "["}) {
			lb = true
			continue
		}
		if lb && tok.Equal(token.Token{Type: token.TYPE_BRACKETCLOSE, Literal: "]"}) {
			break
		}

		if lb {
			title = append(title, tok)
			continue
		}
		url = append(url, tok)
	}

	if len(url) == 0 || !lb {
		tr.Unread(count)
		return ParseResult{}, false
	}
	link := Link{
		URL:   token.Literals(url),
		Title: token.Literals(title),
	}

	return ParseResult{
		Element: link,
		Inner:   []token.Token{},
	}, true
}
