package util

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=string,int"
type ValueType = generic.Type

type ValueTypeSet struct {
	inner map[ValueType]struct{}
}

func NewValueTypeSet() ValueTypeSet {
	return ValueTypeSet{inner: make(map[ValueType]struct{}, 50)}
}

func NewValueTypeSetFromSlice(values []ValueType) ValueTypeSet {
	result := ValueTypeSet{inner: make(map[ValueType]struct{}, 50)}
	for _, v := range values {
		result.Add(v)
	}
	return result
}

func (fr *ValueTypeSet) Clone() ValueTypeSet {
	cloned := make(map[ValueType]struct{}, len(fr.inner))
	for k, v := range fr.inner {
		cloned[k] = v
	}

	return ValueTypeSet{inner: cloned}
}

func (fr *ValueTypeSet) First() ValueType {
	for k, _ := range fr.inner {
		return k
	}

	panic("Empty FieldRange")
}

func (fr *ValueTypeSet) Values() []ValueType {
	values := make([]ValueType, len(fr.inner))
	i := 0
	for val := range fr.inner {
		values[i] = val
		i += 1
	}

	return values
}

func (fr *ValueTypeSet) Len() int {
	return len(fr.inner)
}

func (fr *ValueTypeSet) Add(field ValueType) {
	fr.inner[field] = struct{}{}
}

func (fr *ValueTypeSet) Remove(field ValueType) {
	delete(fr.inner, field)
}

func (fr *ValueTypeSet) Contains(field ValueType) bool {
	_, ok := fr.inner[field]
	return ok
}

func (fr *ValueTypeSet) Intersect(o ValueTypeSet) ValueTypeSet {
	result := NewValueTypeSet()
	for val := range fr.inner {
		if o.Contains(val) {
			result.Add(val)
		}
	}

	return result
}
