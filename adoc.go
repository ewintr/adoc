package adoc

import (
	"time"
)

type ADoc struct {
	Title      string
	Attributes map[string]string
	Author     string
	Path       string
	Date       time.Time
	Content    []Element
}

func NewADoc() *ADoc {
	return &ADoc{
		Attributes: map[string]string{},
		Content:    []Element{},
	}
}
