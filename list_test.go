package adoc_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/go-kit/test"
)

func TestList(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []adoc.Element
	}{
		{
			name:  "one item",
			input: `* item 1`,
			exp: []adoc.Element{
				adoc.List([]adoc.ListItem{
					{adoc.Word("item"), adoc.WhiteSpace(" "), adoc.Word("1")},
				},
				)},
		},
		{
			name: "multiple",
			input: `* item 1
* item 2
* item 3`,
			exp: []adoc.Element{
				adoc.List([]adoc.ListItem{
					{adoc.Word("item"), adoc.WhiteSpace(" "), adoc.Word("1")},
					{adoc.Word("item"), adoc.WhiteSpace(" "), adoc.Word("2")},
					{adoc.Word("item"), adoc.WhiteSpace(" "), adoc.Word("3")},
				})},
		},
		{
			name: "double with pararaph",
			input: `* item 1

* item 2
* item 3

and some text`,
			exp: []adoc.Element{
				adoc.List(
					[]adoc.ListItem{
						{adoc.Word("item"), adoc.WhiteSpace(" "), adoc.Word("1")},
					},
				),
				adoc.List(
					[]adoc.ListItem{
						{adoc.Word("item"), adoc.WhiteSpace(" "), adoc.Word("2")},
						{adoc.Word("item"), adoc.WhiteSpace(" "), adoc.Word("3")},
					},
				),
				adoc.Paragraph([]adoc.Element{
					adoc.Word("and"),
					adoc.WhiteSpace(" "),
					adoc.Word("some"),
					adoc.WhiteSpace(" "),
					adoc.Word("text"),
				}),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := adoc.NewParser(strings.NewReader(tc.input))
			exp := &adoc.ADoc{
				Attributes: map[string]string{},
				Content:    tc.exp,
			}
			test.Equals(t, exp, par.Parse())
		})
	}
}
