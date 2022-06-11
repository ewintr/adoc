package parser_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc/document"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestParser(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *document.Document
	}{
		{
			name: "empty",
			exp:  document.New(),
		},
		{
			name: "codeblock paragraph edge",
			input: `= some title

----
a code block
----

And then some text`,
			exp: &document.Document{
				Title:      "some title",
				Attributes: map[string]string{},
				Content: []element.Element{
					element.CodeBlock{
						element.Word("a"),
						element.WhiteSpace(" "),
						element.Word("code"),
						element.WhiteSpace(" "),
						element.Word("block"),
						element.WhiteSpace("\n"),
					},
					element.Paragraph{
						Elements: []element.Element{
							element.Word("And"),
							element.WhiteSpace(" "),
							element.Word("then"),
							element.WhiteSpace(" "),
							element.Word("some"),
							element.WhiteSpace(" "),
							element.Word("text"),
						},
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := parser.New(strings.NewReader(tc.input))
			test.Equals(t, tc.exp, par.Parse())
		})
	}
}
