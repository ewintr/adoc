package formatter

import (
	"fmt"

	"ewintr.nl/adoc/document"
	"ewintr.nl/adoc/element"
)

type Text struct{}

func NewText() *Text {
	return &Text{}
}

func (t *Text) Format(doc *document.Document) string {
	txt := fmt.Sprintf("%s\n\n", doc.Title)
	txt += t.FormatFragments(doc.Content...)
	return txt
}

func (t *Text) FormatFragments(els ...element.Element) string {
	var text string
	for _, el := range els {
		text += fmt.Sprintf("%s\n\n", el.Text())
	}

	return text
}
