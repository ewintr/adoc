package element_test

import (
	"strings"
	"testing"

	"code.ewintr.nl/adoc/document"
	"code.ewintr.nl/adoc/element"
	"code.ewintr.nl/adoc/parser"
	"code.ewintr.nl/go-kit/test"
)

func TestSubTitle(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []element.Element
	}{
		{
			name:  "empty",
			input: "== ",
			exp:   []element.Element{element.SubTitle("")},
		},
		{
			name:  "subtitle",
			input: "== title with words",
			exp:   []element.Element{element.SubTitle("title with words")},
		},
		{
			name:  "subsubtitle",
			input: "=== title",
			exp:   []element.Element{element.SubSubTitle("title")},
		},
		{
			name:  "trailing newline",
			input: "== title\n",
			exp:   []element.Element{element.SubTitle("title")},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			exp := &document.Document{
				Attributes: map[string]string{},
				Content:    tc.exp,
			}
			par := parser.New(strings.NewReader(tc.input))
			test.Equals(t, exp, par.Parse())
		})
	}
}
