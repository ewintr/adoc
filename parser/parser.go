package parser

import (
	"io"

	"ewintr.nl/adoc"
	"ewintr.nl/adoc/element"
	"ewintr.nl/adoc/token"
)

type Parser struct {
	doc *adoc.ADoc
	tr  *token.TokenReader
	els []element.Element
}

func New(reader io.Reader) *Parser {
	lex := token.NewLexer(reader)
	return NewParserFromChannel(lex.Out())
}

func NewParserFromChannel(toks chan token.Token) *Parser {
	return &Parser{
		doc: adoc.New(),
		tr:  token.NewTokenReader(toks),
		els: []element.Element{},
	}
}

func (p *Parser) Parse() *adoc.ADoc {
	result, ok := element.NewHeaderFromTokens(p.tr)
	if ok {
		if h, ok := result.Element.(element.Header); ok {
			p.doc.Title = h.Title
			p.doc.Author = h.Author
			p.doc.Date = h.Date
			p.doc.Attributes = h.Attributes
		}
	}

	for {
		if done := p.scan(); done {
			break
		}
		p.tr.Discard()
	}

	p.doc.Content = p.els
	return p.doc
}

func (p *Parser) Elements() []element.Element {
	return p.els
}

func (p *Parser) scan() bool {
	if _, ok := p.tr.Read(1); !ok {
		return true
	}
	p.tr.Unread(1)

	type tryFunc func(element.ReadUnreader) (element.ParseResult, bool)
	for _, tf := range []tryFunc{
		element.NewCodeBlockFromTokens,
		element.NewSubTitleFromTokens,
		element.NewListFromTokens,
		element.NewListItemFromTokens,
		element.NewParagraphFromTokens,
		element.NewStyleFromTokens,
		element.NewLinkFromTokens,
	} {
		result, ok := tf(p.tr)
		if !ok {
			continue
		}
		el := result.Element
		if len(result.Inner) != 0 {
			par := NewParserFromChannel(token.NewTokenStream(result.Inner).Out())
			par.Parse()
			el = el.Append(par.Elements())
		}
		p.appendElement(el)
		return false
	}

	p.doPlain()

	return false
}

func (p *Parser) doPlain() {
	tok, ok := p.tr.Read(1)
	if !ok || tok[0].Type == token.TYPE_EOF || tok[0].Type == token.TYPE_EOS {
		return
	}
	p.appendElement(element.MakePlain(tok[0]))
}

func (p *Parser) appendElement(b element.Element) {
	p.els = append(p.els, b)
}
