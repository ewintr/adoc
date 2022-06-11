package formatter

import (
	"ewintr.nl/adoc/document"
	"ewintr.nl/adoc/element"
)

type Formatter interface {
	Format(doc *document.Document) string
	FormatFragments(els ...element.Element) string
}
