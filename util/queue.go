package util

import "container/list"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=string,int"
type ValueTypeQueue struct {
	inner *list.List
}

func NewValueTypeQueue() ValueTypeQueue {
	return ValueTypeQueue{
		inner: list.New(),
	}
}

func (q *ValueTypeQueue) Pop() *ValueType {
	if e := q.inner.Front(); e != nil {
		q.inner.Remove(e)
		cast := e.Value.(ValueType)
		return &cast
	} else {
		return nil
	}
}

func (q *ValueTypeQueue) Peek() *ValueType {
	if e := q.inner.Front(); e != nil {
		cast := e.Value.(ValueType)
		return &cast
	} else {
		return nil
	}
}

func (q *ValueTypeQueue) PushBack(value ValueType) {
	q.inner.PushBack(value)
}
