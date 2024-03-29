package element_test

import (
	"strings"
	"testing"
	"time"

	"code.ewintr.nl/adoc/document"
	"code.ewintr.nl/adoc/element"
	"code.ewintr.nl/adoc/parser"
	"code.ewintr.nl/go-kit/test"
)

func TestHeader(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *document.Document
	}{
		{
			name:  "just title",
			input: "= Title",
			exp: &document.Document{
				Title:      "Title",
				Attributes: map[string]string{},
				Content:    []element.Element{},
			},
		},
		{
			name:  "empty title",
			input: "= ",
			exp:   document.New(),
		},
		{
			name: "full header",
			input: `= Title with words
2022-03-04
Author Name
:key1: value1
:key2: value2

First paragraph`,
			exp: &document.Document{
				Title:  "Title with words",
				Date:   time.Date(2022, time.Month(3), 4, 0, 0, 0, 0, time.UTC),
				Author: "Author Name",
				Attributes: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
				Content: []element.Element{
					element.Paragraph{[]element.Element{
						element.Word("First"),
						element.WhiteSpace(" "),
						element.Word("paragraph"),
					}},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := parser.New(strings.NewReader(tc.input))
			test.Equals(t, tc.exp, par.Parse())
		})
	}
}
