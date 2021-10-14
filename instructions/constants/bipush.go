package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type BIPush struct {
	Value int8
}

func (its *BIPush) FetchOperand(reader *base.BytecodeReader) {
	its.Value = reader.ReadInt8()
}
func (its *BIPush) Execute(frame *rtdaBase.JavaFrame) {
	i := int32(its.Value)
	frame.OperandStack().PushInt(i)
}
