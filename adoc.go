package adoc

import (
	"io"

	"ewintr.nl/adoc/document"
	"ewintr.nl/adoc/formatter"
	"ewintr.nl/adoc/parser"
)

func NewDocument() *document.Document {
	return document.New()
}

func NewParser(reader io.Reader) *parser.Parser {
	return parser.New(reader)
}

func NewTextFormatter() *formatter.Text {
	return formatter.NewText()
}

func NewAsciiDocFormatter() *formatter.AsciiDoc {
	return formatter.NewAsciiDoc()
}

func NewHTMLFormatter() *formatter.HTML {
	return formatter.NewHTML()
}
