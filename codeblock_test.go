package adoc_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
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
				Content:    []adoc.Element{adoc.CodeBlock{}},
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
				Content: []adoc.Element{adoc.CodeBlock{
					adoc.Word("code"),
					adoc.WhiteSpace("\n\n"),
					adoc.Word("more"),
					adoc.WhiteSpace("\n"),
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
				Content: []adoc.Element{
					adoc.Paragraph{
						adoc.Word("----"),
						adoc.WhiteSpace("\n"),
						adoc.Word("code"),
					},
					adoc.Paragraph{
						adoc.Word("more"),
						adoc.WhiteSpace("\n"),
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := adoc.NewParser(strings.NewReader(tc.input))
			test.Equals(t, tc.exp, par.Parse())
		})
	}
}
