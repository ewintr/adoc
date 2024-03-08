package element_test

import (
	"strings"
	"testing"

	"code.ewintr.nl/adoc/document"
	"code.ewintr.nl/adoc/element"
	"code.ewintr.nl/adoc/parser"
	"code.ewintr.nl/go-kit/test"
)

func TestCodeBlock(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *document.Document
	}{
		{
			name: "empty",
			input: `----
----`,
			exp: &document.Document{
				Attributes: map[string]string{},
				Content:    []element.Element{element.CodeBlock{}},
			},
		},
		{
			name: "with newlines",
			input: `----
code

more
----`,
			exp: &document.Document{
				Attributes: map[string]string{},
				Content: []element.Element{element.CodeBlock{
					element.Word("code"),
					element.WhiteSpace("\n\n"),
					element.Word("more"),
					element.WhiteSpace("\n"),
				}},
			},
		},
		{
			name: "with newline at end",
			input: `----
code
----
`,
			exp: &document.Document{
				Attributes: map[string]string{},
				Content: []element.Element{element.CodeBlock{
					element.Word("code"),
					element.WhiteSpace("\n"),
				}},
			},
		},
		{
			name: "missing end",
			input: `----
code

more
`,
			exp: &document.Document{
				Attributes: map[string]string{},
				Content: []element.Element{
					element.Paragraph{[]element.Element{
						element.Word("----"),
						element.WhiteSpace("\n"),
						element.Word("code"),
					}},
					element.Paragraph{[]element.Element{
						element.Word("more"),
						element.WhiteSpace("\n"),
					}},
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
