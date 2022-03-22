package util

type Set[T comparable] struct {
	inner map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{inner: make(map[T]struct{}, 50)}
}

func NewSetFromSlice[T comparable](values []T) Set[T] {
	result := Set[T]{inner: make(map[T]struct{}, 50)}
	for _, v := range values {
		result.Add(v)
	}
	return result
}

func (fr *Set[T]) Clone() Set[T] {
	cloned := make(map[T]struct{}, len(fr.inner))
	for k, v := range fr.inner {
		cloned[k] = v
	}

	return Set[T]{inner: cloned}
}

func (fr *Set[T]) First() T {
	for k, _ := range fr.inner {
		return k
	}

	panic("Empty FieldRange")
}

func (fr *Set[T]) Values() []T {
	values := make([]T, len(fr.inner))
	i := 0
	for val := range fr.inner {
		values[i] = val
		i += 1
	}

	return values
}

func (fr *Set[T]) Len() int {
	return len(fr.inner)
}

func (fr *Set[T]) Add(field T) {
	fr.inner[field] = struct{}{}
}

func (fr *Set[T]) Remove(field T) {
	delete(fr.inner, field)
}

func (fr *Set[T]) Contains(field T) bool {
	_, ok := fr.inner[field]
	return ok
}

func (fr *Set[T]) Intersect(o Set[T]) Set[T] {
	result := NewSet[T]()
	for val := range fr.inner {
		if o.Contains(val) {
			result.Add(val)
		}
	}

	return result
}
