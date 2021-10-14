package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
Operation
	Push null

Format
	aconst_null

Forms
	aconst_null = 1 (0x1)

Operand Stack
	... →
	..., null

Description
	Push the null object reference onto the operand stack.

Notes
	The Java Virtual Machine does not mandate a concrete value for null.

https://docs.oracle.com/javase/specs/jvms/se17/html/jvms-6.html#jvms-6.5.aconst_null
*/
type AConstNull struct {
}

func (its *AConstNull) FetchOperand(reader *base.BytecodeReader) {

}
func (its *AConstNull) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushRef(nil)
}
