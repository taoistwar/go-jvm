package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad struct {
	Index uint
}

func (its *ALoad) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *ALoad) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(uint(its.Index))
	frame.OperandStack().PushRef(value)
}
