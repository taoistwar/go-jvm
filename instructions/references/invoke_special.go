package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type InvokeSpecial struct {
	Index uint
}

func (its *InvokeSpecial) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadUint16())
}

func (its *InvokeSpecial) Execute(frame *rtdaBase.JavaFrame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(its.Index).(*java.MethodRef)
	mylog.Printf("Index:%v, Method:%v%v", its.Index, methodRef.Name(), methodRef.Descriptor())
	frame.OperandStack().PopRef()
}
