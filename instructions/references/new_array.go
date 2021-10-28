package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

const (
	//Array Type  atype(Array type code)
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

/*
Operation
	Create new array

Format
	newarray atype

Forms
	newarray = 188 (0xbc)

Operand Stack
	..., count →
	..., arrayref

Description
	The count must be of type int. It is popped off the operand stack. The count represents the number of elements in the array to be created.

	The atype is a code that indicates the type of array to create.

	A new array whose components are of type atype and of length count is allocated from the garbage-collected heap.
	A reference arrayref to this new array object is pushed into the operand stack.
	Each of the elements of the new array is initialized to the default initial value (§2.3, §2.4) for the element type of the array type.

Run-time Exception
	If count is less than zero, newarray throws a NegativeArraySizeException.

Notes
	In Oracle's Java Virtual Machine implementation, arrays of type boolean (atype is T_BOOLEAN) are stored as arrays of 8-bit values
	and are manipulated using the baload and bastore instructions which also access arrays of type byte.
	Other implementations may implement packed boolean arrays; the baload and bastore instructions must still be used to access those arrays.
*/
type NewArray struct {
	ArrayTypeCode uint8
}

func (its *NewArray) FetchOperand(reader *base.BytecodeReader) {
	its.ArrayTypeCode = reader.ReadUint8()
}

func (its *NewArray) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, its.ArrayTypeCode)
	arr := arrClass.NewJavaArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *java.JavaClassLoader, aType uint8) *java.JavaClass {
	switch aType {
	case AT_BOOLEAN:
		return loader.LoadJClass("[Z")
	case AT_BYTE:
		return loader.LoadJClass("[B")
	case AT_CHAR:
		return loader.LoadJClass("[C")
	case AT_SHORT:
		return loader.LoadJClass("[S")
	case AT_INT:
		return loader.LoadJClass("[I")
	case AT_LONG:
		return loader.LoadJClass("[J")
	case AT_FLOAT:
		return loader.LoadJClass("[F")
	case AT_DOUBLE:
		return loader.LoadJClass("[D")
	default:
		panic("Invalid atype!")
	}
}
