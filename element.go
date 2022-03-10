package adoc

type Element interface {
	Text() string
}

type Empty struct{}

func (e Empty) Text() string { return "" }

type Word string

func (w Word) Text() string { return string(w) }

type WhiteSpace string

func (ws WhiteSpace) Text() string { return string(ws) }

type Image struct {
	Src string
	Alt string
}

func (i Image) Text() string {
	return i.Alt
}
