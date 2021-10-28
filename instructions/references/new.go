package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type BytecodeNew struct {
	Index uint
}

func (its *BytecodeNew) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}

func (its *BytecodeNew) Execute(frame *rtdaBase.JavaFrame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(its.Index)

	switch classRef := classRef.(type) {
	case *java.ClassRef:
		class := classRef.ResolvedClass()
		if class.IsInterface() || class.IsAbstract() {
			panic("java.lang.InstantiationError")
		}
		if !class.InitStarted() {
			frame.RevertNextPC()
			base.InitClass(frame.Thread(), class)
			return
		}
		ref := class.NewObject()
		frame.OperandStack().PushRef(ref)
		mylog.Printf("Index: %v, Class: %v", its.Index, classRef.ClassName())
	default:
		panic("")
	}
}
