package adoc_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/go-kit/test"
)

func TestLink(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []adoc.Element
	}{
		{
			name:  "simple",
			input: "a link[title] somewhere",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Word("a"),
					adoc.WhiteSpace(" "),
					adoc.Link{
						URL:   "link",
						Title: "title",
					},
					adoc.WhiteSpace(" "),
					adoc.Word("somewhere"),
				}),
			},
		},
		{
			name:  "with underscore",
			input: "check https://example.com/some_url[some url]",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Word("check"),
					adoc.WhiteSpace(" "),
					adoc.Link{
						URL:   "https://example.com/some_url",
						Title: "some url",
					},
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
