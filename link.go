package adoc

import "fmt"

type Link struct {
	URL   string
	Title string
}

func (l Link) Text() string { return l.Title }

func (p *Parser) tryLink() bool {
	tok, ok := p.peek()
	if !ok {
		return false
	}
	if tok.Type != TYPE_WORD {
		return false
	}

	url, title := []Token{}, []Token{}
	lb := false
	count := 0
	for {
		tok, ok := p.read()
		if !ok {
			p.unread(count)
			return false
		}
		count++
		if tok.Equal(TOKEN_EOF, TOKEN_NEWLINE, TOKEN_DOUBLE_NEWLINE) {
			p.unread(count)
			return false
		}
		if !lb && tok.Type == TYPE_WHITESPACE {
			p.unread(count)
			return false
		}
		if tok.Equal(Token{Type: TYPE_BRACKETOPEN, Literal: "["}) {
			lb = true
			continue
		}
		if tok.Equal(Token{Type: TYPE_BRACKETCLOSE, Literal: "]"}) {
			break
		}

		if lb {
			title = append(title, tok)
			continue
		}
		url = append(url, tok)
	}
	fmt.Printf("url: %+v, title: %+v\n", url, title)

	if len(url) == 0 || !lb {
		p.unread(count)
		return false
	}
	link := Link{
		URL:   Literals(url),
		Title: Literals(title),
	}
	p.appendElement(link)
	return true
}
