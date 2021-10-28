package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
Operation
	Get length of array

Format
	arraylength

Forms
	arraylength = 190 (0xbe)

Operand Stack
	..., arrayref →
	..., length

Description
	The arrayref must be of type reference and must refer to an array.
	It is popped from the operand stack. The length of the array it references is determined.
	That length is pushed onto the operand stack as an int.

Run-time Exceptions
	If the arrayref is null, the arraylength instruction throws a NullPointerException.
*/
type ArrayLength struct {
}

func (its *ArrayLength) FetchOperand(reader *base.BytecodeReader) {
}

func (its *ArrayLength) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
