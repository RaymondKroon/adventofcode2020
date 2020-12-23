package util

import "fmt"

type String string
type Int int

type SliceType interface {
	Equals(o SliceType) bool
	GreaterThan(o SliceType) bool
	String() string
}

func (s String) Equals(o String) bool {
	return s == o
}

func (s String) GreaterThan(o String) bool {
	return s > o
}

func (s String) String() string {
	return string(s)
}

func (s Int) Equals(o Int) bool {
	return s == o
}

func (s Int) GreaterThan(o Int) bool {
	return s > o
}

func (s Int) String() string {
	return fmt.Sprint(int(s))
}
