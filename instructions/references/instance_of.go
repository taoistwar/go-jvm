package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type InstanceOf struct {
	Index uint
}

func (its *InstanceOf) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}

func (its *InstanceOf) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(its.Index).(*java.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
		mylog.Printf("Index:%v, ThisClass:%v, InstanceOf:%v, Value:%v", its.Index, ref.Class().ThisClassName(), class.ThisClassName(), true)
	} else {
		stack.PushInt(0)
		mylog.Printf("Index:%v, ThisClass:%v, InstanceOf:%v, Value:%v", its.Index, ref.Class().ThisClassName(), class.ThisClassName(), false)
	}

}
