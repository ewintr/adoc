package formatter

import (
	"fmt"
	"html"

	"code.ewintr.nl/adoc/document"
	"code.ewintr.nl/adoc/element"
	"code.ewintr.nl/go-kit/slugify"
)

const htmlPageTemplate = `<!DOCTYPE html>
<html>
<head>
<title>%s</title>
</head>
<body>
%s</body>
</html>
`

type HTML struct{}

func NewHTML() *HTML {
	return &HTML{}
}

func (h *HTML) Format(doc *document.Document) string {
	return fmt.Sprintf(htmlPageTemplate, html.EscapeString(doc.Title), h.FormatFragments(doc.Content...))
}

func (h *HTML) FormatFragments(els ...element.Element) string {
	var html string
	for _, el := range els {
		html += h.htmlElement(el)
	}

	return html
}

func (h *HTML) htmlElement(el element.Element) string {
	switch v := el.(type) {
	case element.SubTitle:
		return fmt.Sprintf("<h2 id=%q>%s</h2>\n", slugify.Slugify(v.Text()), html.EscapeString(v.Text()))
	case element.SubSubTitle:
		return fmt.Sprintf("<h3 id=%q>%s</h3>\n", slugify.Slugify(v.Text()), html.EscapeString(v.Text()))
	case element.List:
		var items []element.Element
		for _, i := range v {
			items = append(items, i)
		}
		return fmt.Sprintf("<ul>\n%s</ul>\n", h.FormatFragments(items...))
	case element.ListItem:
		return fmt.Sprintf("<li>%s</li>\n", h.FormatFragments(v...))
	case element.CodeBlock:
		return fmt.Sprintf("<pre><code>%s</code></pre>", html.EscapeString(v.Text()))
	case element.Paragraph:
		return fmt.Sprintf("<p>%s</p>\n", h.FormatFragments(v.Elements...))
	case element.Strong:
		return fmt.Sprintf("<strong>%s</strong>", h.FormatFragments(v...))
	case element.Emphasis:
		return fmt.Sprintf("<em>%s</em>", h.FormatFragments(v...))
	case element.Code:
		return fmt.Sprintf("<code>%s</code>", h.FormatFragments(v...))
	case element.Link:
		return fmt.Sprintf("<a href=%q>%s</a>", v.URL, html.EscapeString(v.Title))
	case element.Word:
		return html.EscapeString(v.Text())
	case element.WhiteSpace:
		return " "
	default:
		return ""
	}
}
