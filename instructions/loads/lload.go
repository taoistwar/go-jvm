package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LLoad struct {
	Index uint
}

func (its *LLoad) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *LLoad) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(uint(its.Index))
	frame.OperandStack().PushInt(value)
}
