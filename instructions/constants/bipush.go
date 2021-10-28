package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
Operation
	Push byte
Format:
	bipush byte
Forms:
	bipush = 16 (0x10)
Operand Stack
	... →
	..., value
Description
	The immediate byte is sign-extended to an int value. That value is pushed onto the operand stack.
*/
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
