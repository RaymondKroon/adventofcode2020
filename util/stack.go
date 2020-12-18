package util

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ValueType=string,int"
type ValueTypeStack struct {
	inner []ValueType
}

func NewValueTypeStack() ValueTypeStack {
	return ValueTypeStack{
		inner: make([]ValueType, 0, 10),
	}
}

func (stack *ValueTypeStack) Pop() *ValueType {
	if len(stack.inner) == 0 {
		return nil
	} else {
		lastN := len(stack.inner) - 1
		pop := stack.inner[lastN]
		stack.inner = stack.inner[:lastN:lastN]
		return &pop
	}
}

func (stack *ValueTypeStack) Peek() *ValueType {
	if len(stack.inner) == 0 {
		return nil
	} else {
		return &stack.inner[len(stack.inner)-1]
	}
}

func (stack *ValueTypeStack) Push(value ValueType) {
	stack.inner = append(stack.inner, value)
}

func (stack *ValueTypeStack) ToArray() []ValueType {
	return stack.inner
}
