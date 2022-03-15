package element_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestCodeBlock(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *adoc.ADoc
	}{
		{
			name: "empty",
			input: `----
----`,
			exp: &adoc.ADoc{
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
			exp: &adoc.ADoc{
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
			name: "missing end",
			input: `----
code

more
`,
			exp: &adoc.ADoc{
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
