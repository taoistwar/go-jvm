package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type SIPush struct {
	value int16
}

func (its *SIPush) FetchOperand(reader *base.BytecodeReader) {
	its.value = reader.ReadInt16()
}
func (its *SIPush) Execute(frame *rtdaBase.JavaFrame) {
	i := int32(its.value)
	frame.OperandStack().PushInt(i)
}
