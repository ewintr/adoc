package element_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc/document"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestParagraph(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *document.Document
	}{
		{
			name:  "single paragraph",
			input: "some text",
			exp: &document.Document{
				Attributes: map[string]string{},
				Content: []element.Element{
					element.Paragraph{Elements: []element.Element{
						element.Word("some"),
						element.WhiteSpace(" "),
						element.Word("text"),
					}},
				}},
		},
		{
			name: "title with paragraphs",
			input: `= Title

paragraph one

paragraph two`,
			exp: &document.Document{
				Title:      "Title",
				Attributes: map[string]string{},
				Content: []element.Element{
					element.Paragraph{Elements: []element.Element{
						element.Word("paragraph"),
						element.WhiteSpace(" "),
						element.Word("one"),
					}},
					element.Paragraph{Elements: []element.Element{
						element.Word("paragraph"),
						element.WhiteSpace(" "),
						element.Word("two"),
					}},
				},
			},
		},
		{
			name: "three with trailing newline",
			input: `one

two

three
`,
			exp: &document.Document{
				Attributes: map[string]string{},
				Content: []element.Element{
					element.Paragraph{Elements: []element.Element{element.Word("one")}},
					element.Paragraph{Elements: []element.Element{element.Word("two")}},
					element.Paragraph{Elements: []element.Element{element.Word("three"), element.WhiteSpace("\n")}},
				}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := parser.New(strings.NewReader(tc.input))
			test.Equals(t, tc.exp, par.Parse())
		})
	}
}
