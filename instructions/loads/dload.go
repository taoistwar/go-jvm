package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DLoad struct {
	Index uint
}

func (its *DLoad) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *DLoad) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetDouble(uint(its.Index))
	frame.OperandStack().PushDouble(value)
}
