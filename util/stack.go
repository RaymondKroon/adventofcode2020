package util

type Stack[T any] struct {
	inner []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		inner: make([]T, 0, 10),
	}
}

func (stack *Stack[T]) Pop() *T {
	if len(stack.inner) == 0 {
		return nil
	} else {
		lastN := len(stack.inner) - 1
		pop := stack.inner[lastN]
		stack.inner = stack.inner[:lastN:lastN]
		return &pop
	}
}

func (stack *Stack[T]) Peek() *T {
	if len(stack.inner) == 0 {
		return nil
	} else {
		return &stack.inner[len(stack.inner)-1]
	}
}

func (stack *Stack[T]) Push(value T) {
	stack.inner = append(stack.inner, value)
}

func (stack *Stack[T]) ToArray() []T {
	return stack.inner
}
