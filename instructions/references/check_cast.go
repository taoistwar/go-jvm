package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type CheckCast struct {
	Index uint
}

func (its *CheckCast) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}

func (its *CheckCast) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(its.Index).(*java.ClassRef)
	class := classRef.ResolvedClass()
	mylog.Printf("Index: %v, Object: %v IsInstanceOf Class: %v", its.Index, ref.Class().ThisClassName(), class.ThisClassName())
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
