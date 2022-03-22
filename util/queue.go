package util

import "container/list"

type Queue[T any] struct {
	inner *list.List
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		inner: list.New(),
	}
}

func NewQueueFromSlice[T any](slice []T) Queue[T] {
	queue := NewQueue[T]()
	for _, e := range slice {
		queue.PushBack(e)
	}

	return queue
}

func (q *Queue[T]) Len() int {
	return q.inner.Len()
}

func (q *Queue[T]) Pop() *T {
	if e := q.inner.Front(); e != nil {
		q.inner.Remove(e)
		cast := e.Value.(T)
		return &cast
	} else {
		return nil
	}
}

func (q *Queue[T]) Peek() *T {
	if e := q.inner.Front(); e != nil {
		cast := e.Value.(T)
		return &cast
	} else {
		return nil
	}
}

func (q *Queue[T]) PushBack(value T) {
	q.inner.PushBack(value)
}

func (q *Queue[T]) Values() []T {
	values := make([]T, 0, q.inner.Len())
	val := q.inner.Front()
	for val != nil {
		values = append(values, val.Value.(T))
		val = val.Next()
	}
	return values
}
