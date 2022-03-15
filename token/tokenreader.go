package token

type TokenReader struct {
	in chan Token
	ts []Token
	r  int
}

func NewTokenReader(in chan Token) *TokenReader {
	return &TokenReader{
		in: in,
		ts: []Token{},
	}
}

func (tr *TokenReader) Read(n int) ([]Token, bool) {
	toks := []Token{}
	for i := 0; i < n; i++ {
		tok, ok := tr.readOne()
		if !ok {
			tr.Unread(len(toks))
			return []Token{}, false
		}
		toks = append(toks, tok)
	}
	return toks, true
}

func (tr *TokenReader) Unread(n int) bool {
	//if n > tr.r {
	//	return false
	//}
	for i := n; i > 0; i-- {
		tr.r--
	}
	return true
}

func (tr *TokenReader) Discard() {
	tr.ts = tr.ts[tr.r:]
	tr.r = 0
}

func (tr *TokenReader) readOne() (Token, bool) {
	if tr.r == len(tr.ts) {
		if ok := tr.consume(); !ok {
			return Token{}, false
		}
	}

	tok := tr.ts[tr.r]
	tr.r++

	return tok, true
}

func (tr *TokenReader) consume() bool {
	tok, ok := <-tr.in
	if !ok {
		return false
	}
	tr.ts = append(tr.ts, tok)

	return true
}
