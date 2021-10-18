package base

import "github.com/taoistwar/go-jvm/rtda/java"

/*
	一个线程包含一个线程栈，线程栈包含一组栈帧，每次方法调用产生一个栈帧。
	一个栈帧包括局部变量表、操作数栈。
*/
type JavaThread struct {
	pc    int
	stack *JavaStack
}

func NewJavaThread() *JavaThread {
	return &JavaThread{stack: newJavaStack(1024)}
}

func (thread *JavaThread) PC() int {
	return thread.pc
}

func (thread *JavaThread) SetPC(pc int) {
	thread.pc = pc
}

func (thread *JavaThread) PushFrame(frame *JavaFrame) {
	thread.stack.push(frame)
}

func (thread *JavaThread) PopFrame() *JavaFrame {
	return thread.stack.pop()
}

func (thread *JavaThread) TopFrame() *JavaFrame {
	return thread.stack.top
}

func (thread *JavaThread) CurrentFrame() *JavaFrame {
	return thread.stack.current()
}

func (thread *JavaThread) NewJavaFrame(javaMethod *java.JavaMethod) *JavaFrame {
	return NewJavaFrame(thread, javaMethod)
}
