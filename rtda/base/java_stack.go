package base

type JavaStack struct {
	maxSize uint // 线程堆大小
	size    uint
	top     *JavaFrame
}

func newJavaStack(maxSize uint) *JavaStack {
	return &JavaStack{maxSize: maxSize}
}

func (stack *JavaStack) push(frame *JavaFrame) {
	if stack.size >= stack.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if stack.top != nil {
		frame.lower = stack.top
	}
	stack.top = frame
	stack.size++
}

func (stack *JavaStack) pop() *JavaFrame {
	if stack.top == nil {
		panic("jvm stack is empty")
	}
	top := stack.top
	stack.top = top.lower
	top.lower = nil
	stack.size--
	return top
}

func (stack *JavaStack) current() *JavaFrame {
	if stack.top == nil {
		panic("jvm stack is empty")
	}
	return stack.top
}
