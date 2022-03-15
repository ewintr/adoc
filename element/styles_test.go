package element_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestStyles(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []element.Element
	}{
		{
			name:  "strong",
			input: "*strong*",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Strong{element.Word("strong")},
				},
				},
			},
		},
		{
			name:  "emphasis",
			input: "_emphasis_",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Emphasis{element.Word("emphasis")},
				},
				},
			},
		},
		{
			name:  "code",
			input: "`code`",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Code{element.Word("code")},
				},
				},
			},
		},
		{
			name:  "mixed",
			input: "some `code code` in plain",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Word("some"),
					element.WhiteSpace(" "),
					element.Code{
						element.Word("code"),
						element.WhiteSpace(" "),
						element.Word("code"),
					},
					element.WhiteSpace(" "),
					element.Word("in"),
					element.WhiteSpace(" "),
					element.Word("plain"),
				},
				},
			},
		},
		{
			name:  "incomplete",
			input: "a *word",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Word("a"),
					element.WhiteSpace(" "),
					element.Word("*"),
					element.Word("word"),
				}},
			},
		},
		{
			name:  "trailing space",
			input: "*word *",
			exp: []element.Element{
				element.Paragraph{Elements: []element.Element{
					element.Word("*"),
					element.Word("word"),
					element.WhiteSpace(" "),
					element.Word("*"),
				}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := parser.New(strings.NewReader(tc.input))
			exp := &adoc.ADoc{
				Attributes: map[string]string{},
				Content:    tc.exp,
			}
			test.Equals(t, exp, par.Parse())
		})
	}
}
