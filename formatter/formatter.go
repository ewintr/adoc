package formatter

import (
	"code.ewintr.nl/adoc/document"
	"code.ewintr.nl/adoc/element"
)

type Formatter interface {
	Format(doc *document.Document) string
	FormatFragments(els ...element.Element) string
}
