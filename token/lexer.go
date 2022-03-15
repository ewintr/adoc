package token

import (
	"bufio"
	"bytes"
	"io"
)

type Lexer struct {
	in  *bufio.Reader
	out chan Token
	err error
}

func NewLexer(in io.Reader) *Lexer {
	lex := &Lexer{
		in:  bufio.NewReader(in),
		out: make(chan Token),
	}

	go lex.run()

	return lex
}

func (lex *Lexer) Out() chan Token {
	return lex.out
}

func (lex *Lexer) Error() error {
	return lex.err
}

func (lex *Lexer) run() {
	for {
		tok := lex.scan()
		if tok.Type == TYPE_EOF {
			lex.out <- TOKEN_EOF
			close(lex.out)
			return
		}

		lex.out <- tok
	}
}

func (lex *Lexer) scan() Token {
	var buf bytes.Buffer

	ch := lex.read()
	if _, err := buf.WriteRune(ch); err != nil {
		lex.err = err
		return Token{Type: TYPE_EOF}
	}

	tt := FindTokenType(ch)

LEXSCANLOOP:
	for {
		ch = lex.read()
		switch {
		case ch == eof:
			break LEXSCANLOOP
		case FindTokenType(ch) != tt:
			lex.unread()
			break LEXSCANLOOP
		default:
			if _, err := buf.WriteRune(ch); err != nil {
				lex.err = err
				return Token{Type: TYPE_EOF}
			}
		}
	}

	return Token{
		Type:    tt,
		Literal: buf.String(),
	}
}

func (lex *Lexer) read() rune {
	ch, _, err := lex.in.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (lex *Lexer) unread() {
	lex.in.UnreadRune()
}

func FindTokenType(ch rune) TokenType {
	for tt, rr := range TokenTypeMap {
		for _, r := range rr {
			if ch == r {
				return tt
			}
		}
	}

	return TYPE_WORD
}
