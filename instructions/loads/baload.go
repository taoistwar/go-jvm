package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
Operation
	Load byte or boolean from array

Format
	baload

Forms
	baload = 51 (0x33)

Operand Stack
	..., arrayref, index →
	..., value

Description
	The arrayref must be of type reference and must refer to an array whose components are of type byte or of type boolean.
	The index must be of type int.
	Both arrayref and index are popped from the operand stack.
	The byte value in the component of the array at index is retrieved, sign-extended to an int value, and pushed onto the top of the operand stack.

Run-time Exceptions
	If arrayref is null, baload throws a NullPointerException.

	Otherwise, if index is not within the bounds of the array referenced by arrayref, the baload instruction throws an ArrayIndexOutOfBoundsException.

Notes
	The baload instruction is used to load values from both byte and boolean arrays.
	In Oracle's Java Virtual Machine implementation, boolean arrays - that is, arrays of type T_BOOLEAN (§2.2, §newarray)
	 - are implemented as arrays of 8-bit values.
	Other implementations may implement packed boolean arrays; the baload instruction of such implementations must be used to access those arrays.

*/
type BaLoad struct {
}

func (its *BaLoad) FetchOperand(reader *base.BytecodeReader) {
}

func (its *BaLoad) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Bytes()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}
