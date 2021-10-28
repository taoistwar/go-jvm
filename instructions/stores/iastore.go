package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
Operation
	Store into int array

Format
	iastore

Forms
	iastore = 79 (0x4f)

Operand Stack
	..., arrayref, index, value →
	...

Description
	The arrayref must be of type reference and must refer to an array whose components are of type int. Both index and value must be of type int. The arrayref, index, and value are popped from the operand stack. The int value is stored as the component of the array indexed by index.

Run-time Exceptions
	If arrayref is null, iastore throws a NullPointerException.

	Otherwise, if index is not within the bounds of the array referenced by arrayref, the iastore instruction throws an ArrayIndexOutOfBoundsException.


*/
type IAStore struct {
}

func (its *IAStore) FetchOperand(reader *base.BytecodeReader) {
}

func (its *IAStore) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}
