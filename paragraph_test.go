package adoc_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/go-kit/test"
)

func TestParagraph(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *adoc.ADoc
	}{
		{
			name:  "single paragraph",
			input: "some text",
			exp: &adoc.ADoc{
				Attributes: map[string]string{},
				Content: []adoc.Element{
					adoc.Paragraph([]adoc.Element{
						adoc.Word("some"),
						adoc.WhiteSpace(" "),
						adoc.Word("text"),
					}),
				}},
		},
		{
			name: "title with paragraphs",
			input: `= Title

paragraph one

paragraph two`,
			exp: &adoc.ADoc{
				Title:      "Title",
				Attributes: map[string]string{},
				Content: []adoc.Element{
					adoc.Paragraph([]adoc.Element{
						adoc.Word("paragraph"),
						adoc.WhiteSpace(" "),
						adoc.Word("one"),
					}),
					adoc.Paragraph([]adoc.Element{
						adoc.Word("paragraph"),
						adoc.WhiteSpace(" "),
						adoc.Word("two"),
					}),
				},
			},
		},
		{
			name: "three with trailing newline",
			input: `one

two

three
`,
			exp: &adoc.ADoc{
				Attributes: map[string]string{},
				Content: []adoc.Element{
					adoc.Paragraph([]adoc.Element{adoc.Word("one")}),
					adoc.Paragraph([]adoc.Element{adoc.Word("two")}),
					adoc.Paragraph([]adoc.Element{adoc.Word("three"), adoc.WhiteSpace("\n")}),
				}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := adoc.NewParser(strings.NewReader(tc.input))
			test.Equals(t, tc.exp, par.Parse())
		})
	}
}
