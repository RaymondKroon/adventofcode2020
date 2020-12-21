package util

type String string
type Int int

type SliceType interface {
	Equals(o SliceType) bool
}

func (s String) Equals(o String) bool {
	return s == o
}

func (s Int) Equals(o Int) bool {
	return s == o
}
