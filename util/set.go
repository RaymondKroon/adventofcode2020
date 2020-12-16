package util

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=string,int"
type ValueType = generic.Type

type ValueTypeSet struct {
	inner map[ValueType]bool
}

func NewValueTypeSet() ValueTypeSet {
	return ValueTypeSet{inner: make(map[ValueType]bool, 50)}
}

func (fr *ValueTypeSet) First() ValueType {
	for k, _ := range fr.inner {
		return k
	}

	panic("Empty FieldRange")
}

func (fr *ValueTypeSet) Len() int {
	return len(fr.inner)
}

func (fr *ValueTypeSet) Add(field ValueType) {
	fr.inner[field] = true
}

func (fr *ValueTypeSet) Remove(field ValueType) {
	delete(fr.inner, field)
}
