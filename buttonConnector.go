package goco

import "honnef.co/go/js/dom"

type ButtonConnector struct {
	Connector

	domNode *dom.HTMLButtonElement
}

func NewButtonConnector() *ButtonConnector {
	return &ButtonConnector{}
}

func (b *ButtonConnector) Attach(selector string) {
}

func (b *ButtonConnector) Detach() {}
