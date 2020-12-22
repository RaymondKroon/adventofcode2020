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

func NewValueTypeQueueFromSlice(slice []ValueType) ValueTypeQueue {
	queue := NewValueTypeQueue()
	for _, e := range slice {
		queue.PushBack(e)
	}

	return queue
}

func (q *ValueTypeQueue) Len() int {
	return q.inner.Len()
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

func (q *ValueTypeQueue) Values() []ValueType {
	values := make([]ValueType, 0, q.inner.Len())
	val := q.inner.Front()
	for val != nil {
		values = append(values, val.Value.(ValueType))
		val = val.Next()
	}
	return values
}
