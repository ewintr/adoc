package adoc_test

import (
	"strings"
	"testing"
	"time"

	"ewintr.nl/adoc"
	"ewintr.nl/go-kit/test"
)

func TestHeader(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *adoc.ADoc
	}{
		{
			name:  "just title",
			input: "= Title",
			exp: &adoc.ADoc{
				Title:      "Title",
				Attributes: map[string]string{},
				Content:    []adoc.Element{},
			},
		},
		{
			name:  "empty title",
			input: "= ",
			exp:   adoc.NewADoc(),
		},
		{
			name: "full header",
			input: `= Title with words
2022-03-04
Author Name
:key1: value1
:key2: value2

First paragraph`,
			exp: &adoc.ADoc{
				Title:  "Title with words",
				Date:   time.Date(2022, time.Month(3), 4, 0, 0, 0, 0, time.UTC),
				Author: "Author Name",
				Attributes: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				Content: []adoc.Element{
					adoc.Paragraph([]adoc.Element{
						adoc.Word("First"),
						adoc.WhiteSpace(" "),
						adoc.Word("paragraph"),
					}),
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := adoc.NewParser(strings.NewReader(tc.input))
			test.Equals(t, tc.exp, par.Parse())
		})
	}
}
