package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
Operation
	Do nothing

Format
	nop

Forms
	nop = 0 (0x0)

Operand Stack
	No change

Description
	Do nothing.

https://docs.oracle.com/javase/specs/jvms/se17/html/jvms-6.html#jvms-6.5.nop
*/
type NOP struct {
}

func (its *NOP) FetchOperand(reader *base.BytecodeReader) {

}
func (its *NOP) Execute(frame *rtdaBase.JavaFrame) {

}
