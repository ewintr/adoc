package token_test

import (
	"strings"
	"testing"
	"time"

	"ewintr.nl/adoc/token"
	"ewintr.nl/go-kit/test"
)

func TestLexer(t *testing.T) {
	word := token.TYPE_WORD
	ws := token.TYPE_WHITESPACE
	nl := token.TYPE_NEWLINE
	eq := token.TYPE_EQUALSIGN
	bt := token.TYPE_BACKTICK
	as := token.TYPE_ASTERISK
	un := token.TYPE_UNDERSCORE

	for _, tc := range []struct {
		name  string
		input string
		exp   []token.Token
	}{
		{
			name:  "word string",
			input: "one two",
			exp: []token.Token{
				{Type: word, Literal: "one"},
				{Type: ws, Literal: " "},
				{Type: word, Literal: "two"},
			},
		},
		{
			name:  "punctuation",
			input: `. ,`,
			exp: []token.Token{
				{Type: word, Literal: "."},
				{Type: ws, Literal: " "},
				{Type: word, Literal: ","},
			},
		},
		{
			name:  "whitespace",
			input: " \t",
			exp: []token.Token{
				{Type: ws, Literal: " \t"},
			},
		},
		{
			name:  "tab",
			input: "\t",
			exp:   []token.Token{{Type: ws, Literal: "\t"}},
		},
		{
			name:  "newlines",
			input: "\n\n\n",
			exp:   []token.Token{{Type: nl, Literal: "\n\n\n"}},
		},
		{
			name:  "special chars",
			input: "=*_",
			exp: []token.Token{
				{Type: eq, Literal: "="},
				{Type: as, Literal: "*"},
				{Type: un, Literal: "_"},
			},
		},
		{
			name:  "mixed",
			input: "This is a line with mixed \t `stuff`, see\t==?",
			exp: []token.Token{
				{Type: word, Literal: "This"},
				{Type: ws, Literal: " "},
				{Type: word, Literal: "is"},
				{Type: ws, Literal: " "},
				{Type: word, Literal: "a"},
				{Type: ws, Literal: " "},
				{Type: word, Literal: "line"},
				{Type: ws, Literal: " "},
				{Type: word, Literal: "with"},
				{Type: ws, Literal: " "},
				{Type: word, Literal: "mixed"},
				{Type: ws, Literal: " \t "},
				{Type: bt, Literal: "`"},
				{Type: word, Literal: "stuff"},
				{Type: bt, Literal: "`"},
				{Type: word, Literal: ","},
				{Type: ws, Literal: " "},
				{Type: word, Literal: "see"},
				{Type: ws, Literal: "\t"},
				{Type: eq, Literal: "=="},
				{Type: word, Literal: "?"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			lex := token.NewLexer(input)
			act := []token.Token{}
			stop := time.Now().Add(3 * time.Second)

		T:
			for {
				select {
				case tok, ok := <-lex.Out():
					if !ok {
						break T
					}
					act = append(act, tok)
				default:
					if time.Now().After(stop) {
						break T
					}
					time.Sleep(5 * time.Millisecond)
				}
			}

			test.OK(t, lex.Error())
			exp := append(tc.exp, token.TOKEN_EOF)
			test.Equals(t, exp, act)
		})
	}
}
