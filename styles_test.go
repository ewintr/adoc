package adoc_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/go-kit/test"
)

func TestStyles(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   []adoc.Element
	}{
		{
			name:  "strong",
			input: "*strong*",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Strong{adoc.Word("strong")},
				},
				),
			},
		},
		{
			name:  "emphasis",
			input: "_emphasis_",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Emphasis{adoc.Word("emphasis")},
				},
				),
			},
		},
		{
			name:  "code",
			input: "`code`",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Code{adoc.Word("code")},
				},
				),
			},
		},
		{
			name:  "mixed",
			input: "some `code code` in plain",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Word("some"),
					adoc.WhiteSpace(" "),
					adoc.Code{
						adoc.Word("code"),
						adoc.WhiteSpace(" "),
						adoc.Word("code"),
					},
					adoc.WhiteSpace(" "),
					adoc.Word("in"),
					adoc.WhiteSpace(" "),
					adoc.Word("plain"),
				},
				),
			},
		},
		{
			name:  "incomplete",
			input: "a *word",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Word("a"),
					adoc.WhiteSpace(" "),
					adoc.Word("*"),
					adoc.Word("word"),
				}),
			},
		},
		{
			name:  "trailing space",
			input: "*word *",
			exp: []adoc.Element{
				adoc.Paragraph([]adoc.Element{
					adoc.Word("*"),
					adoc.Word("word"),
					adoc.WhiteSpace(" "),
					adoc.Word("*"),
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
