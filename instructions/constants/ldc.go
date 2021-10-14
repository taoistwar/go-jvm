package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*

 */
type Ldc struct {
	Index uint
}

func (its *Ldc) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}
func (its *Ldc) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(its.Index)

	switch c := c.(type) {
	case int32:
		stack.PushInt(c)
		mylog.Printf("Index: %v, Value: %v", its.Index, c)
	case float32:
		stack.PushFloat(c)
		mylog.Printf("Index: %v, Value: %v", its.Index, c)
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}
