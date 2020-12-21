// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package util

import "container/list"

type StringQueue struct {
	inner *list.List
}

func NewStringQueue() StringQueue {
	return StringQueue{
		inner: list.New(),
	}
}

func (q *StringQueue) Pop() *string {
	if e := q.inner.Front(); e != nil {
		q.inner.Remove(e)
		cast := e.Value.(string)
		return &cast
	} else {
		return nil
	}
}

func (q *StringQueue) Peek() *string {
	if e := q.inner.Front(); e != nil {
		cast := e.Value.(string)
		return &cast
	} else {
		return nil
	}
}

func (q *StringQueue) PushBack(value string) {
	q.inner.PushBack(value)
}

type IntQueue struct {
	inner *list.List
}

func NewIntQueue() IntQueue {
	return IntQueue{
		inner: list.New(),
	}
}

func (q *IntQueue) Pop() *int {
	if e := q.inner.Front(); e != nil {
		q.inner.Remove(e)
		cast := e.Value.(int)
		return &cast
	} else {
		return nil
	}
}

func (q *IntQueue) Peek() *int {
	if e := q.inner.Front(); e != nil {
		cast := e.Value.(int)
		return &cast
	} else {
		return nil
	}
}

func (q *IntQueue) PushBack(value int) {
	q.inner.PushBack(value)
}