package element

import (
	"time"

	"ewintr.nl/adoc/token"
)

type Header struct {
	Title      string
	Author     string
	Date       time.Time
	Attributes map[string]string
}

func (h Header) Text() string               { return h.Title }
func (h Header) Append(_ []Element) Element { return h }

func NewHeaderFromTokens(tr ReadUnreader) (ParseResult, bool) {
	toks, ok := tr.Read(2)
	if !ok {
		return ParseResult{}, false
	}
	if !toks[0].Equal(token.Token{Type: token.TYPE_EQUALSIGN, Literal: "="}) {
		tr.Unread(2)
		return ParseResult{}, false
	}
	if toks[1].Type != token.TYPE_WHITESPACE {
		tr.Unread(2)
		return ParseResult{}, false
	}
	for {
		ntoks, ok := tr.Read(1)
		if !ok {
			tr.Unread(len(toks))
			return ParseResult{}, false
		}
		if ntoks[0].Equal(token.TOKEN_EOF, token.TOKEN_DOUBLE_NEWLINE) {
			break
		}
		toks = append(toks, ntoks[0])
	}

	h := &Header{
		Attributes: map[string]string{},
	}
	lines := token.Split(toks[2:], token.TYPE_NEWLINE)
	h.Title = token.Literals(lines[0])

	for _, line := range lines[1:] {
		switch {
		case tryHeaderDate(h, line):
			continue
		case tryHeaderField(h, line):
			continue
		default:
			h.Author = token.Literals(line)
		}
	}

	return ParseResult{
		Element: *h,
		Inner:   []token.Token{},
	}, true
}

func tryHeaderField(h *Header, line []token.Token) bool {
	if len(line) < 4 {
		return false
	}
	pair := token.Split(line, token.TYPE_WHITESPACE)
	if len(pair) != 2 {
		return false
	}
	key, value := pair[0], pair[1]

	if !token.HasPattern(key, []token.TokenType{token.TYPE_COLON, token.TYPE_WORD, token.TYPE_COLON}) {
		return false
	}

	h.Attributes[key[1].Literal] = token.Literals(value)

	return true
}

func tryHeaderDate(h *Header, line []token.Token) bool {
	date, err := time.Parse("2006-01-02", token.Literals(line))
	if err != nil {
		return false
	}
	h.Date = date

	return true
}
