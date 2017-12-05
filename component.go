package goco

import "honnef.co/go/js/dom"

// "honnef.co/go/js/console"

// type Component interface {
// 	ViewModel() *js.M
// 	Template() string
// }

type Component interface {
	Element() dom.Element
}

type ObservableString interface {
	// New(s string) ObservableString
	Set(s string)
	Get() string
	AddBefore(do func()) int
	RemoveBefore(id int)
	AddAfter(do func())
	RemoveAfter(id int)
}

type BeforeHandler struct {
	id  int
	fun func(string, string)
}

type AfterHandler struct {
	id  int
	fun func(string)
}

type BeforeHandlerMap map[int]*BeforeHandler

type AfterHandlerMap map[int]*AfterHandler

type Observed struct {
	lastBeforeID   int
	lastAfterID    int
	beforeHandlers BeforeHandlerMap
	afterHandlers  AfterHandlerMap
}

type Selectable interface {
	QuerySelector(string) dom.Element
}

type Stylable interface {
	SetStyle(string, string)
	GetStyle(string) string
}

// func SetStyle(s Stylable, name string, value string) {
// 	s.rootElement.Style().SetProperty(name, value, "")
// }

func CreateHTMLElement(markup string) dom.HTMLElement {
	el := dom.GetWindow().Document().CreateElement("div")
	el.SetInnerHTML(markup)
	// for now, just hope that the element is an HTMLElement --
	// in theory it could be anything, but since the layout is
	// defined in this package, it isn't.
	return el.FirstElementChild().(dom.HTMLElement)
}
