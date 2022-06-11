package document

import (
	"time"

	"ewintr.nl/adoc/element"
)

type Document struct {
	Title      string
	Attributes map[string]string
	Author     string
	Date       time.Time
	Content    []element.Element
}

func New() *Document {
	return &Document{
		Attributes: map[string]string{},
		Content:    []element.Element{},
	}
}
