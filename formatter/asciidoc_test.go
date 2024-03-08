package formatter_test

import (
	"testing"
	"time"

	"code.ewintr.nl/adoc/document"
	"code.ewintr.nl/adoc/element"
	"code.ewintr.nl/adoc/formatter"
	"code.ewintr.nl/go-kit/test"
)

func TestAsciiDoc(t *testing.T) {
	input := &document.Document{
		Title:  "A Title",
		Author: "Author",
		Date:   time.Date(2022, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
		Attributes: map[string]string{
			"key1": "value 1",
			"key2": "value 2",
		},
		Content: []element.Element{
			element.Paragraph{
				Elements: []element.Element{
					element.Word("some"),
					element.WhiteSpace(" "),
					element.Word("text"),
				},
			},
		},
	}

	exp := `= A Title
Author
2022-06-11
:key1: value 1
:key2: value 2

some text

`

	test.Equals(t, exp, formatter.NewAsciiDoc().Format(input))
}

func TestAsciiDocFragment(t *testing.T) {
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
			name: "pararaphs",
			input: element.Paragraph{
				Elements: []element.Element{
					element.Word("a"),
					element.WhiteSpace(" "),
					element.Word("word"),
				},
			},
			exp: "a word\n\n",
		},
		{
			name: "strong",
			input: element.Strong{
				element.Word("something"),
				element.WhiteSpace(" "),
				element.Word("strong"),
			},
			exp: "*something strong*",
		},
		{
			name: "nested",
			input: element.Paragraph{
				Elements: []element.Element{
					element.Word("normal"),
					element.WhiteSpace(" "),
					element.Word("text"),
					element.WhiteSpace(" "),
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
			exp: "normal text * and strong* too\n\n",
		},
		{
			name: "emphasis",
			input: element.Emphasis{
				element.Word("yes"),
			},
			exp: "_yes_",
		},
		{
			name: "code",
			input: element.Code{
				element.Word("simple"),
			},
			exp: "`simple`",
		},
		{
			name: "link",
			input: element.Link{
				URL:   "http://example.com",
				Title: "an example",
			},
			exp: `http://example.com[an example]`,
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
			exp: `* item 1
* item 2

`,
		},
		{
			name: "code block",
			input: element.CodeBlock{
				element.Word("some"),
				element.WhiteSpace(" "),
				element.Word("text"),
				element.WhiteSpace("\n"),
				element.Word("<p>with</p>"),
				element.WhiteSpace("\t"),
				element.Word("formatting"),
			},
			exp: `----
some text
<p>with</p>	formatting
----

`,
		},
		{
			name:  "subtitle",
			input: element.SubTitle("a subtitle"),
			exp:   "== a subtitle\n\n",
		},
		{
			name:  "subsubtitle",
			input: element.SubSubTitle("a subsubtitle"),
			exp:   "=== a subsubtitle\n\n",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			test.Equals(t, tc.exp, formatter.NewAsciiDoc().FormatFragments(tc.input))
		})
	}
}
