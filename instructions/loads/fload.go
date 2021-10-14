package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FLoad struct {
	Index uint
}

func (its *FLoad) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *FLoad) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetFloat(uint(its.Index))
	frame.OperandStack().PushFloat(value)
}
