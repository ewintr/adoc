package formatter

import (
	"fmt"

	"code.ewintr.nl/adoc/document"
	"code.ewintr.nl/adoc/element"
)

type AsciiDoc struct{}

func NewAsciiDoc() *AsciiDoc {
	return &AsciiDoc{}
}

func (ad *AsciiDoc) Format(doc *document.Document) string {
	return fmt.Sprintf("%s\n%s", asciiDocHeader(doc), ad.FormatFragments(doc.Content...))
}

func asciiDocHeader(doc *document.Document) string {
	header := fmt.Sprintf("= %s\n", doc.Title)
	if doc.Author != "" {
		header += fmt.Sprintf("%s\n", doc.Author)
	}
	if !doc.Date.IsZero() {
		header += fmt.Sprintf("%s\n", doc.Date.Format("2006-01-02"))
	}
	for k, v := range doc.Attributes {
		header += fmt.Sprintf(":%s: %s\n", k, v)
	}

	return header
}

func (ad *AsciiDoc) FormatFragments(els ...element.Element) string {
	var asciiDoc string
	for _, el := range els {
		asciiDoc += ad.asciiDocElement(el)
	}

	return asciiDoc
}

func (ad *AsciiDoc) asciiDocElement(el element.Element) string {
	switch v := el.(type) {
	case element.SubTitle:
		return fmt.Sprintf("== %s\n\n", v.Text())
	case element.SubSubTitle:
		return fmt.Sprintf("=== %s\n\n", v.Text())
	case element.List:
		var items []element.Element
		for _, i := range v {
			items = append(items, i)
		}
		return fmt.Sprintf("%s\n", ad.FormatFragments(items...))
	case element.ListItem:
		return fmt.Sprintf("* %s\n", ad.FormatFragments(v...))
	case element.CodeBlock:
		return fmt.Sprintf("----\n%s\n----\n\n", v.Text())
	case element.Paragraph:
		return fmt.Sprintf("%s\n\n", ad.FormatFragments(v.Elements...))
	case element.Strong:
		return fmt.Sprintf("*%s*", ad.FormatFragments(v...))
	case element.Emphasis:
		return fmt.Sprintf("_%s_", ad.FormatFragments(v...))
	case element.Code:
		return fmt.Sprintf("`%s`", ad.FormatFragments(v...))
	case element.Link:
		return fmt.Sprintf("%s[%s]", v.URL, v.Title)
	case element.Word:
		return v.Text()
	case element.WhiteSpace:
		return " "
	default:
		return ""
	}
}
