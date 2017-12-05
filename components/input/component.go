package input

import (
	"errors"
	"fmt"
	"log"

	"github.com/adapterik/goco"
	"honnef.co/go/js/dom"
)

type Input struct {
	rootElement dom.HTMLElement

	inputElement *dom.HTMLInputElement

	value *goco.ObservedString
}

func New() *Input {
	el := goco.CreateHTMLElement(T_markup())
	self := &Input{
		rootElement:  el,
		inputElement: el.(*dom.HTMLInputElement),
		value:        goco.NewObservedString()}
	self.Attach()
	return self

}

func (i *Input) Element() dom.Element {
	return i.rootElement
}

// func New(parent *dom.HTMLDivElement) *Component {
// 	self := &Component{parentNode: parent, value: goco.NewObservedString()}
// 	self.Attach()
// 	return self
// }

func (i *Input) SetStyle(name string, value string) {
	i.rootElement.Style().SetProperty(name, value, "")
}

func selectDivElement(from goco.Selectable, selector string) (*dom.HTMLDivElement, error) {
	n := from.QuerySelector(selector)

	if n == nil {
		return nil, errors.New("Sorry could not find dom node")
	}
	switch n.(type) {
	default:
		return nil, fmt.Errorf("Not an Element! %s", selector)
	case *dom.HTMLDivElement:
		return n.(*dom.HTMLDivElement), nil
	}
}

func selectInputElement(from goco.Selectable, selector string) (*dom.HTMLInputElement, error) {
	n := from.QuerySelector(selector)

	if n == nil {
		return nil, errors.New("Sorry could not find dom node")
	}
	switch n.(type) {
	default:
		return nil, fmt.Errorf("Not an Input Element! %s", selector)
	case *dom.HTMLInputElement:
		return n.(*dom.HTMLInputElement), nil
	}
}

func selectElement(from goco.Selectable, selector string) (dom.Element, error) {
	n := from.QuerySelector(selector)

	if n == nil {
		return nil, errors.New("Sorry could not find dom node")
	}

	return n, nil
}

// Attach this component to a parent node identified by the given
// selector string.
// TODO: This needs a dom context, or more simply a node.
func (i *Input) Attach() error {
	// set up listeners for this input, this is the magic bit which
	// provides the mutation binding.
	// We use the listeners to update our string observable.
	i.inputElement.AddEventListener("change", false, func(ev dom.Event) {
		v := ev.Target().(*dom.HTMLInputElement).Value
		i.value.Set(v)
	})

	i.inputElement.AddEventListener("keyup", false, func(ev dom.Event) {
		v := ev.Target().(*dom.HTMLInputElement).Value
		i.value.Set(v)
	})

	return nil
}

func (i *Input) Detach() {
}

// setDomValue is a convenience function for setting the
// nodes value.
func (i *Input) setDomValue(s string) {
	if i.inputElement == nil {
		panic("Not connected to root node")
	}
	i.inputElement.Value = s
}

// SetValue sets the string observable, and also
// sets the dom value itself.
func (i *Input) SetValue(s string) {
	log.Println("setting dom node value to ", s)
	i.value.Set(s)
	i.setDomValue(s)
}

// GetValue gets the value from the string observable,
// not directly from the node's value. The node's value
// _should_ be the same as the observable.
func (i *Input) GetValue() string {
	if i.inputElement == nil {
		panic("Not connected")
	}
	return i.value.Get()
	// return i.rootNode.Value
}

// OnValue allows a user of this component to react to
// changes in the string observable.
func (i *Input) OnValue(fn func(string)) {
	i.value.AddAfter(fn)
}
