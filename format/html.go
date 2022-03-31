package format

import (
	"fmt"
	"html"

	"ewintr.nl/adoc"
	"ewintr.nl/adoc/element"
	"ewintr.nl/go-kit/slugify"
)

const pageTemplate = `<!DOCTYPE html>
<html>
<head>
<title>%s</title>
</head>
<body>
%s</body>
</html>
`

func HTML(doc *adoc.ADoc) string {
	return fmt.Sprintf(pageTemplate, html.EscapeString(doc.Title), HTMLFragment(doc.Content...))
}

func HTMLFragment(els ...element.Element) string {
	var html string
	for _, el := range els {
		html += htmlElement(el)
	}

	return html
}

func htmlElement(el element.Element) string {
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
		return fmt.Sprintf("<ul>\n%s</ul>\n", HTMLFragment(items...))
	case element.ListItem:
		return fmt.Sprintf("<li>%s</li>\n", HTMLFragment(v...))
	case element.CodeBlock:
		return fmt.Sprintf("<pre><code>%s</code></pre>", html.EscapeString(v.Text()))
	case element.Paragraph:
		return fmt.Sprintf("<p>%s</p>\n", HTMLFragment(v.Elements...))
	case element.Strong:
		return fmt.Sprintf("<strong>%s</strong>", HTMLFragment(v...))
	case element.Emphasis:
		return fmt.Sprintf("<em>%s</em>", HTMLFragment(v...))
	case element.Code:
		return fmt.Sprintf("<code>%s</code>", HTMLFragment(v...))
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
