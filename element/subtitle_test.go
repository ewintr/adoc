package element_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
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
			exp := &adoc.ADoc{
				Attributes: map[string]string{},
				Content:    tc.exp,
			}
			par := parser.New(strings.NewReader(tc.input))
			test.Equals(t, exp, par.Parse())
		})
	}
}
