package parser_test

import (
	"strings"
	"testing"

	"ewintr.nl/adoc"
	"ewintr.nl/adoc/parser"
	"ewintr.nl/go-kit/test"
)

func TestParser(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		exp   *adoc.ADoc
	}{
		{
			name: "empty",
			exp:  adoc.NewADoc(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			par := parser.New(strings.NewReader(tc.input))
			test.Equals(t, tc.exp, par.Parse())
		})
	}
}
