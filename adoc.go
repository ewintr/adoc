package adoc

import (
	"time"

	"ewintr.nl/adoc/element"
)

type ADoc struct {
	Title      string
	Attributes map[string]string
	Author     string
	Path       string
	Date       time.Time
	Content    []element.Element
}

func NewADoc() *ADoc {
	return &ADoc{
		Attributes: map[string]string{},
		Content:    []element.Element{},
	}
}
