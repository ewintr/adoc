package formatter_test

import (
	"strings"
	"testing"

	"code.ewintr.nl/adoc/formatter"
	"code.ewintr.nl/adoc/parser"
	"code.ewintr.nl/go-kit/test"
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
