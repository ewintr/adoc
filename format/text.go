package format

import (
	"fmt"

	"ewintr.nl/adoc"
)

func Text(doc *adoc.ADoc) string {
	txt := fmt.Sprintf("%s\n\n", doc.Title)
	for _, el := range doc.Content {
		txt += fmt.Sprintf("%s\n\n", el.Text())
	}

	return txt
}
