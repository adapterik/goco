package text

import (
	"github.com/adapterik/goco"
	"honnef.co/go/js/dom"
)

type Component struct {
	// parentNode *dom.HTMLDivElement
	rootNode *dom.HTMLDivElement
}

func New() *Component {
	el := goco.CreateHTMLElement(T_markup())
	self := &Component{rootNode: el.(*dom.HTMLDivElement)}
	return self
}

func (i *Component) Element() dom.Element {
	return i.rootNode
}

func (i *Component) SetStyle(name string, value string) {
	i.rootNode.Style().SetProperty(name, value, "")
}

// Attach this component to a parent node identified by the given
// selector string.
// TODO: This needs a dom context, or more simply a node.
// func (i *Component) Attach() error {
// 	i.parentNode.SetInnerHTML(T_markup())
// 	i.rootNode = i.parentNode.QuerySelector("div").(*dom.HTMLDivElement)
// 	return nil
// }

func (i *Component) Detach() {
}

func (i *Component) SetText(s string) {
	if i.rootNode == nil {
		panic("Not connected")
	}
	i.rootNode.SetTextContent(s)
}

func (i *Component) GetText() string {
	if i.rootNode == nil {
		panic("Not connected")
	}
	return i.rootNode.TextContent()
}
