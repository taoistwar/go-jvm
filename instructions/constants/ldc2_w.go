package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
ldc系列指令从运行时常量池中加载常量值，并把它推入操作数栈。
*/
type Ldc2W struct {
	Index uint
}

func (its *Ldc2W) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}
func (its *Ldc2W) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	constant := cp.GetConstant(its.Index)

	switch constant := constant.(type) {
	case int64:
		stack.PushLong(constant)
	case float64:
		stack.PushDouble(constant)
	default:
		panic("java.lang.ClassFormatError")
	}
}
