// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package util

type StringStack struct {
	inner []string
}

func NewStringStack() StringStack {
	return StringStack{
		inner: make([]string, 0, 10),
	}
}

func (stack *StringStack) Pop() *string {
	if len(stack.inner) == 0 {
		return nil
	} else {
		lastN := len(stack.inner) - 1
		pop := stack.inner[lastN]
		stack.inner = stack.inner[:lastN:lastN]
		return &pop
	}
}

func (stack *StringStack) Peek() *string {
	if len(stack.inner) == 0 {
		return nil
	} else {
		return &stack.inner[len(stack.inner)-1]
	}
}

func (stack *StringStack) Push(value string) {
	stack.inner = append(stack.inner, value)
}

func (stack *StringStack) ToArray() []string {
	return stack.inner
}

type IntStack struct {
	inner []int
}

func NewIntStack() IntStack {
	return IntStack{
		inner: make([]int, 0, 10),
	}
}

func (stack *IntStack) Pop() *int {
	if len(stack.inner) == 0 {
		return nil
	} else {
		lastN := len(stack.inner) - 1
		pop := stack.inner[lastN]
		stack.inner = stack.inner[:lastN:lastN]
		return &pop
	}
}

func (stack *IntStack) Peek() *int {
	if len(stack.inner) == 0 {
		return nil
	} else {
		return &stack.inner[len(stack.inner)-1]
	}
}

func (stack *IntStack) Push(value int) {
	stack.inner = append(stack.inner, value)
}

func (stack *IntStack) ToArray() []int {
	return stack.inner
}
