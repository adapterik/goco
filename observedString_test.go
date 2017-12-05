package goco

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	s := NewObservedString()
	s.AddBefore(func(_ string, _ string) { fmt.Println("about to set..") })
	s.AddAfter(func(_ string) { fmt.Println("just set...") })
	expected := "Hi!"
	s.Set(expected)
	actual := s.Get()
	if expected != actual {
		t.Errorf("Sorry, it didn't work")
	}
}

func Test2(t *testing.T) {
	s := NewObservedString()
	expected := "hi"
	var actual string
	s.AddBefore(func(_ string, _ string) {
		fmt.Println("here")
		actual = expected
	})
	s.Set("hello")
	if expected != actual {
		t.Errorf("Sorry, it didn't work")
	}
}

func Test3(t *testing.T) {
	s := NewObservedString()
	expected := "hi"
	var actual string
	s.AddAfter(func(_ string) {
		fmt.Println("here")
		actual = expected
	})
	s.Set("hello")
	if expected != actual {
		t.Errorf("Sorry, it didn't work")
	}
}

func TestOneTalksToAnother(t *testing.T) {
	s1 := NewObservedString()
	s2 := NewObservedString()

	expected := "hi"
	var actual string

	s1.AddAfter(func(newValue string) {
		s2.Set(newValue)
	})
	s2.AddAfter(func(newValue string) {
		actual = newValue
	})
	s1.Set(expected)
	if expected != actual {
		t.Errorf("Sorry, it didn't work")
	}
}
