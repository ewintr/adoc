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

func New() *ADoc {
	return &ADoc{
		Attributes: map[string]string{},
		Content:    []element.Element{},
	}
}
