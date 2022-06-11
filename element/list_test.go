package element_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc/document"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestList(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []element.Element
	}{
		{
			name:  "one item",
			input: `* item 1`,
			exp: []element.Element{
				element.List([]element.ListItem{
					{element.Word("item"), element.WhiteSpace(" "), element.Word("1")},
				},
				)},
		},
		{
			name: "multiple",
			input: `* item 1
* item 2
* item 3`,
			exp: []element.Element{
				element.List([]element.ListItem{
					{element.Word("item"), element.WhiteSpace(" "), element.Word("1")},
					{element.Word("item"), element.WhiteSpace(" "), element.Word("2")},
					{element.Word("item"), element.WhiteSpace(" "), element.Word("3")},
				})},
		},
		{
			name: "double with pararaph",
			input: `* item 1

* item 2
* item 3

and some text`,
			exp: []element.Element{
				element.List(
					[]element.ListItem{
						{element.Word("item"), element.WhiteSpace(" "), element.Word("1")},
					},
				),
				element.List(
					[]element.ListItem{
						{element.Word("item"), element.WhiteSpace(" "), element.Word("2")},
						{element.Word("item"), element.WhiteSpace(" "), element.Word("3")},
					},
				),
				element.Paragraph{Elements: []element.Element{
					element.Word("and"),
					element.WhiteSpace(" "),
					element.Word("some"),
					element.WhiteSpace(" "),
					element.Word("text"),
				}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := parser.New(strings.NewReader(tc.input))
			exp := &document.Document{
				Attributes: map[string]string{},
				Content:    tc.exp,
			}
			test.Equals(t, exp, par.Parse())
		})
	}
}
