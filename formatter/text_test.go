package formatter_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc/formatter"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestText(t *testing.T) {
	input := `= A Title

Some document

With some text.`

	exp := `A Title

Some document

With some text.

`

	doc := parser.New(strings.NewReader(input)).Parse()
	test.Equals(t, exp, formatter.NewText().Format(doc))
}
