package base

import (
	"github.com/taoistwar/go-jvm/rtda/java"
)

type JavaFrame struct {
	lower        *JavaFrame
	localVars    JavaLocalVars
	operandStack *JavaOperandStack
	thread       *JavaThread
	method       *java.JavaMethod
	nextPC       int // the next instruction after the call
}

func NewJavaFrame(thread *JavaThread, javaMethod *java.JavaMethod) *JavaFrame {
	return &JavaFrame{
		localVars:    newJavaLocalVars(javaMethod.MaxLocals()),
		operandStack: newJavaOperandStack(javaMethod.MaxStack()),
		thread:       thread,
		method:       javaMethod,
		nextPC:       0,
	}
}

func (frame *JavaFrame) OperandStack() *JavaOperandStack {
	return frame.operandStack
}

func (frame *JavaFrame) LocalVars() *JavaLocalVars {
	return &frame.localVars
}

func (frame *JavaFrame) NextPC() int {
	return frame.nextPC
}
func (frame *JavaFrame) Thread() *JavaThread {
	return frame.thread
}
func (frame *JavaFrame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}
func (frame *JavaFrame) Method() *java.JavaMethod {
	return frame.method
}
