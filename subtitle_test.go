package adoc_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/go-kit/test"
)

func TestSubTitle(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []adoc.Element
	}{
		{
			name:  "empty",
			input: "== ",
			exp:   []adoc.Element{adoc.SubTitle("")},
		},
		{
			name:  "subtitle",
			input: "== title with words",
			exp:   []adoc.Element{adoc.SubTitle("title with words")},
		},
		{
			name:  "subsubtitle",
			input: "=== title",
			exp:   []adoc.Element{adoc.SubSubTitle("title")},
		},
		{
			name:  "trailing newline",
			input: "== title\n",
			exp:   []adoc.Element{adoc.SubTitle("title")},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			exp := &adoc.ADoc{
				Attributes: map[string]string{},
				Content:    tc.exp,
			}
			par := adoc.NewParser(strings.NewReader(tc.input))
			test.Equals(t, exp, par.Parse())
		})
	}
}
