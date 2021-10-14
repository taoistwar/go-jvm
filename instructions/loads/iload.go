package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ILoad struct {
	Index uint
}

func (its *ILoad) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *ILoad) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(uint(its.Index))
	frame.OperandStack().PushInt(value)
}
