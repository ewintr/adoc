package format_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc/format"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestText(t *testing.T) {
	input := `= A Title

Some Document

With some text`
	exp := `A Title

Some Document

With some text

`

	doc := parser.New(strings.NewReader(input)).Parse()
	test.Equals(t, exp, format.Text(doc))
}
