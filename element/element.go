package element

import "ewintr.nl/adoc/token"

type Element interface {
	Text() string
	Append([]Element) Element
}

type ParseResult struct {
	Element Element
	Inner   []token.Token
}

type ReadUnreader interface {
	Read(n int) ([]token.Token, bool)
	Unread(n int) bool
}

type Empty struct{}

func (e Empty) Text() string { return "" }

type Image struct {
	Src string
	Alt string
}

func (i Image) Text() string {
	return i.Alt
}
