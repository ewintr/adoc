package token

const (
	TYPE_EOF TokenType = iota
	TYPE_EOS
	TYPE_WHITESPACE
	TYPE_NEWLINE

	TYPE_EQUALSIGN
	TYPE_ASTERISK
	TYPE_UNDERSCORE
	TYPE_BACKTICK
	TYPE_DASH
	TYPE_COLON
	TYPE_BRACKETOPEN
	TYPE_BRACKETCLOSE

	TYPE_WORD
)

var (
	eof = rune(0)
)

var TokenTypeMap = map[TokenType][]rune{
	TYPE_EOF:          {eof},
	TYPE_WHITESPACE:   {' ', '\t'},
	TYPE_NEWLINE:      {'\n'},
	TYPE_EQUALSIGN:    {'='},
	TYPE_BACKTICK:     {'`'},
	TYPE_DASH:         {'-'},
	TYPE_COLON:        {':'},
	TYPE_ASTERISK:     {'*'},
	TYPE_UNDERSCORE:   {'_'},
	TYPE_BRACKETOPEN:  {'['},
	TYPE_BRACKETCLOSE: {']'},
}

type TokenType int

var (
	TOKEN_EOF            = Token{Type: TYPE_EOF}
	TOKEN_EOS            = Token{Type: TYPE_EOS}
	TOKEN_NEWLINE        = Token{Type: TYPE_NEWLINE, Literal: "\n"}
	TOKEN_DOUBLE_NEWLINE = Token{Type: TYPE_NEWLINE, Literal: "\n\n"}
	TOKEN_ASTERISK       = Token{Type: TYPE_ASTERISK, Literal: "*"}
)

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) Len() int {
	return len(t.Literal)
}

func (t Token) Equal(wanted ...Token) bool {
	for _, w := range wanted {
		if t.Type == w.Type && t.Literal == w.Literal {
			return true
		}
	}

	return false
}

func Literals(ts []Token) string {
	var s string
	for _, t := range ts {
		s += t.Literal
	}

	return s
}

func StartsWith(ts []Token, pat []Token) bool {
	if len(ts) < len(pat) {
		return false
	}
	for i, t := range ts {
		if !t.Equal(pat[i]) {
			return false
		}
	}

	return true
}

func EndsWith(ts []Token, pat []Token) bool {
	if len(ts) < len(pat) {
		return false
	}
	for i := 1; i <= len(pat); i++ {
		if !ts[len(ts)-i].Equal(pat[len(pat)-i]) {
			return false
		}
	}

	return true
}

func HasPattern(ts []Token, pat []TokenType) bool {
	if len(ts) != len(pat) {
		return false
	}
	for i, t := range ts {
		if t.Type != pat[i] {
			return false
		}
	}

	return true
}

func Split(ts []Token, tt TokenType) [][]Token {
	lines := [][]Token{}
	line := []Token{}
	for _, t := range ts {
		if t.Type == tt {
			lines = append(lines, line)
			line = []Token{}
			continue
		}
		line = append(line, t)
	}
	lines = append(lines, line)

	return lines
}

func SplitOnFirst(ts []Token, tt TokenType) ([]Token, []Token) {
	if len(ts) < 2 {
		return ts, []Token{}
	}

	first := ts[0].Type
	for i, tok := range ts[1:] {
		if tok.Type == first {
			return ts[1:i], ts[i:]
		}
	}

	return ts, []Token{}
}
