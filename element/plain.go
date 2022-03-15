package element

import "ewintr.nl/adoc/token"

type Word string

func (w Word) Text() string               { return string(w) }
func (w Word) Append(_ []Element) Element { return w }

type WhiteSpace string

func (ws WhiteSpace) Text() string               { return string(ws) }
func (ws WhiteSpace) Append(_ []Element) Element { return ws }

func MakePlain(tok token.Token) Element {
	switch tok.Type {
	case token.TYPE_WHITESPACE:
		return WhiteSpace(tok.Literal)
	case token.TYPE_NEWLINE:
		return WhiteSpace(tok.Literal)
	default:
		return Word(tok.Literal)
	}

}
