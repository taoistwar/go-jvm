package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type InvokeStatic struct {
	Index uint
}

func (its *InvokeStatic) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}

func (its *InvokeStatic) Execute(frame *rtdaBase.JavaFrame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(its.Index).(*java.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	mylog.Printf("Index: %v, Class: %v, Method: %v %v", its.Index, resolvedMethod.Class().ThisClassName(), resolvedMethod.Name(), resolvedMethod.Descriptor())

	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolvedMethod.Class()
	// 类初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	// 方法调用
	base.InvokeMethod(frame, resolvedMethod)
}
