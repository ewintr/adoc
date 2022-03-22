package format_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/format"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestHTML(t *testing.T) {
	input := `= A Title

Some document

With some text.`

	exp := `<!DOCTYPE html>
<html>
<head>
<title>A Title</title>
</head>
<body>
<p>Some document</p>
<p>With some text.</p>
</body>
</html>
`
	doc := parser.New(strings.NewReader(input)).Parse()
	test.Equals(t, exp, format.HTML(doc))
}

func TestHTMLFragment(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input element.Element
		exp   string
	}{
		{
			name:  "whitespace",
			input: element.WhiteSpace("\n"),
			exp:   " ",
		},
		{
			name:  "word",
			input: element.Word("word"),
			exp:   "word",
		},
		{
			name:  "word with html",
			input: element.Word("<h1>hi</h1>"),
			exp:   "&lt;h1&gt;hi&lt;/h1&gt;",
		},
		{
			name: "pararaphs",
			input: element.Paragraph{
				Elements: []element.Element{
					element.Word("a"),
					element.WhiteSpace(" "),
					element.Word("word"),
				},
			},
			exp: "<p>a word</p>\n",
		},
		{
			name: "strong",
			input: element.Strong{
				element.Word("something"),
				element.WhiteSpace(" "),
				element.Word("strong"),
			},
			exp: "<strong>something strong</strong>",
		},
		{
			name: "nested",
			input: element.Paragraph{
				Elements: []element.Element{
					element.Word("normal"),
					element.WhiteSpace(" "),
					element.Word("text"),
					element.Strong{
						element.WhiteSpace(" "),
						element.Word("and"),
						element.WhiteSpace(" "),
						element.Word("strong"),
					},
					element.WhiteSpace(" "),
					element.Word("too"),
				},
			},
			exp: "<p>normal text<strong> and strong</strong> too</p>\n",
		},
		{
			name: "emphasis",
			input: element.Emphasis{
				element.Word("yes"),
			},
			exp: "<em>yes</em>",
		},
		{
			name: "code",
			input: element.Code{
				element.Word("simple"),
			},
			exp: "<code>simple</code>",
		},
		{
			name: "link",
			input: element.Link{
				URL:   "http://example.com",
				Title: "an example",
			},
			exp: `<a href="http://example.com">an example</a>`,
		},
		{
			name: "list",
			input: element.List{
				element.ListItem{
					element.Word("item"),
					element.WhiteSpace(" "),
					element.Word("1"),
				},
				element.ListItem{
					element.Word("item"),
					element.WhiteSpace(" "),
					element.Word("2"),
				},
			},
			exp: `<ul>
<li>item 1</li>
<li>item 2</li>
</ul>
`,
		},
		{
			name: "code block",
			input: element.CodeBlock{
				element.Word("some"),
				element.WhiteSpace(" "),
				element.Word("text"),
				element.WhiteSpace("\n"),
				element.Word("with"),
				element.WhiteSpace("\t"),
				element.Word("formatting"),
			},
			exp: `<pre><code>some text
with	formatting</code></pre>`,
		},
		{
			name:  "subtitle",
			input: element.SubTitle("a subtitle"),
			exp:   "<h2 id=\"a-subtitle\">a subtitle</h2>\n",
		},
		{
			name:  "subsubtitle",
			input: element.SubSubTitle("a subsubtitle"),
			exp:   "<h3 id=\"a-subsubtitle\">a subsubtitle</h3>\n",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, format.HTMLFragment(tc.input))
		})
	}
}
