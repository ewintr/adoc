package format

import (
	"fmt"

	"ewintr.nl/adoc"
	"ewintr.nl/adoc/element"
)

func AsciiDoc(doc *adoc.ADoc) string {
	return fmt.Sprintf("%s\n%s", AsciiDocHeader(doc), AsciiDocFragment(doc.Content...))
}

func AsciiDocHeader(doc *adoc.ADoc) string {
	header := fmt.Sprintf("= %s\n", doc.Title)
	for k, v := range doc.Attributes {
		header += fmt.Sprintf(":%s: %s\n", k, v)
	}

	return header
}

func AsciiDocFragment(els ...element.Element) string {
	var asciiDoc string
	for _, el := range els {
		asciiDoc += asciiDocElement(el)
	}

	return asciiDoc
}

func asciiDocElement(el element.Element) string {
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
		return fmt.Sprintf("%s\n", AsciiDocFragment(items...))
	case element.ListItem:
		return fmt.Sprintf("* %s\n", AsciiDocFragment(v...))
	case element.CodeBlock:
		return fmt.Sprintf("----\n%s\n----\n\n", v.Text())
	case element.Paragraph:
		return fmt.Sprintf("%s\n\n", AsciiDocFragment(v.Elements...))
	case element.Strong:
		return fmt.Sprintf("*%s*", AsciiDocFragment(v...))
	case element.Emphasis:
		return fmt.Sprintf("_%s_", AsciiDocFragment(v...))
	case element.Code:
		return fmt.Sprintf("`%s`", AsciiDocFragment(v...))
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
