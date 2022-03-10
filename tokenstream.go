package adoc

type TokenStream struct {
	ts  []Token
	out chan Token
}

func NewTokenStream(toks []Token) *TokenStream {
	stream := &TokenStream{
		ts:  toks,
		out: make(chan Token),
	}

	go stream.run()

	return stream
}

func (s *TokenStream) run() {
	for _, tok := range s.ts {
		s.out <- tok
	}
	close(s.out)
}

func (s *TokenStream) Out() chan Token {
	return s.out
}
