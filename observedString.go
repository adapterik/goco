package goco

import "fmt"

type ObservedString struct {
	Observed
	currentValue string
	pastValues   []string
}

func NewObservedString() *ObservedString {
	return &ObservedString{Observed: Observed{beforeHandlers: make(BeforeHandlerMap), afterHandlers: make(AfterHandlerMap)}}
}

func (s *ObservedString) runBefore(newValue string) {
	for k, h := range s.beforeHandlers {
		fmt.Printf("Running before handler %d\n", k)
		h.fun(s.currentValue, newValue)
	}
}

func (s *ObservedString) runAfter() {
	for k, h := range s.afterHandlers {
		fmt.Printf("Running after handler %d\n", k)
		h.fun(s.currentValue)
	}
}

func (s *ObservedString) Set(newValue string) {
	// call befores
	s.runBefore(newValue)
	s.currentValue = newValue
	s.runAfter()
	// call afters
}

func (s *ObservedString) Get() string {
	return s.currentValue
}

func (s *ObservedString) AddBefore(f func(string, string)) {
	s.lastBeforeID++
	h := BeforeHandler{id: s.lastBeforeID, fun: f}
	s.beforeHandlers[s.lastBeforeID] = &h
}

func (s *ObservedString) AddAfter(f func(string)) {
	s.lastAfterID++
	h := AfterHandler{id: s.lastAfterID, fun: f}
	s.afterHandlers[s.lastAfterID] = &h
}

func (s *ObservedString) RemoveBefore(idToRemove int) {
	delete(s.beforeHandlers, idToRemove)
}

func (s *ObservedString) RemoveAfter(idToRemove int) {
	delete(s.afterHandlers, idToRemove)
}
