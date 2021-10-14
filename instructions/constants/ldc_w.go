package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*

 */
type LdcW struct {
	Index uint
}

func (its *LdcW) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}
func (its *LdcW) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(its.Index)

	switch c := c.(type) {
	case int32:
		stack.PushInt(c)
	case float32:
		stack.PushFloat(c)
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}
