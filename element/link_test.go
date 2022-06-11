package element_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc/document"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestLink(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []element.Element
	}{
		{
			name:  "simple",
			input: "a link[title] somewhere",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Word("a"),
					element.WhiteSpace(" "),
					element.Link{
						URL:   "link",
						Title: "title",
					},
					element.WhiteSpace(" "),
					element.Word("somewhere"),
				}},
			},
		},
		{
			name:  "with underscore",
			input: "check https://example.com/some_url[some url]",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Word("check"),
					element.WhiteSpace(" "),
					element.Link{
						URL:   "https://example.com/some_url",
						Title: "some url",
					},
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
