package common

import (
	"github.com/richardzhang0301/isokit"
	"honnef.co/go/js/dom"
)

type Env struct {
	TemplateSet    *isokit.TemplateSet
	Window         dom.Window
	Document       dom.Document
	PrimaryContent dom.Element
}
