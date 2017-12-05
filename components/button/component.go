package button

import (
	"github.com/adapterik/goco"
	"honnef.co/go/js/dom"
)

type Button struct {
	rootElement dom.HTMLElement
	// eventHandlers EventHandlers
	clickListener *dom.EventListenerID
}

type EventHandler struct {
	fn        func(dom.Event)
	eventType string
	capture   bool
	id        *dom.EventListenerID
}

type EventHandlers map[string]*EventHandler

func NewButton() *Button {
	el := goco.CreateHTMLElement(T_markup())
	// ehs := make(EventHandlers)
	self := &Button{
		rootElement: el.(dom.HTMLElement)}
	// eventHandlers: ehs}
	return self
}

func (b *Button) Element() dom.Element {
	return b.rootElement
}

func (b *Button) SetLabel(label string) {
	b.rootElement.SetTextContent(label)
}

func (b *Button) OnClick(fn func(ev dom.Event)) {
	// eh := EventHandler{
	// 	fn:        fn,
	// 	eventType: "click",
	// 	capture:   false}
	// id := uuid.NewV4().String()
	// b.eventHandlers[id] = &eh
	b.clickListener = b.rootElement.AddEventListener2("click", false, fn)

}

func (b *Button) OffClick(id string) {
	// eh := b.eventHandlers[id]
	// if eh == nil {
	// 	return
	// }
	b.rootElement.RemoveEventListener2(b.clickListener)
}

// Various styles that may be set.

func (b *Button) SetBackgroundColor(color string) {
	b.rootElement.Style().SetProperty("background-color", color, "")
}

func (b *Button) SetHoverBackgroundColor(color string) {
	b.rootElement.Style().SetProperty("background-color", color, "")
}

func (b *Button) SetColor(color string) {
	b.rootElement.Style().SetProperty("color", color, "")
}
