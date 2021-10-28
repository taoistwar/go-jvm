package references

import (
	"fmt"

	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

/*
Operation
	Create new multidimensional array

Format
	multianewarray indexbyte1 indexbyte2 dimensions

Forms
	multianewarray = 197 (0xc5)

Operand Stack
	..., count1, [count2, ...] →
	..., arrayref

Description
	The dimensions operand is an unsigned byte that must be greater than or equal to 1.
	It represents the number of dimensions of the array to be created. The operand stack must contain dimensions values.
	Each such value represents the number of components in a dimension of the array to be created, must be of type int,
	and must be non-negative. The count1 is the desired length in the first dimension, count2 in the second, etc.

	All of the count values are popped off the operand stack. The unsigned indexbyte1 and indexbyte2 are used to
	construct an index into the run-time constant pool of the current class (§2.6),
	where the value of the index is (indexbyte1 << 8) | indexbyte2.
	The run-time constant pool item at the index must be a symbolic reference to a class, array, or interface type.
	The named class, array, or interface type is resolved (§5.4.3.1).
	The resulting entry must be an array class type of dimensionality greater than or equal to dimensions.

	A new multidimensional array of the array type is allocated from the garbage-collected heap.
	If any count value is zero, no subsequent dimensions are allocated.
	The components of the array in the first dimension are initialized to subarrays of the type of the second dimension,
	and so on. The components of the last allocated dimension of the array are initialized to the default initial value
	 (§2.3, §2.4) for the element type of the array type.
	A reference arrayref to the new array is pushed onto the operand stack.

Linking Exceptions
	During resolution of the symbolic reference to the class, array, or interface type, any of the exceptions documented in §5.4.3.1 can be thrown.

	Otherwise, if the current class does not have permission to access the element type of the resolved array class, multianewarray throws an IllegalAccessError.

Run-time Exception
	Otherwise, if any of the dimensions values on the operand stack are less than zero, the multianewarray instruction throws a NegativeArraySizeException.

Notes
	It may be more efficient to use newarray or anewarray (§newarray, §anewarray) when creating an array of a single dimension.

The array class referenced via the run-time constant pool may have more dimensions than the dimensions operand of the multianewarray instruction. In that case, only the first dimensions of the dimensions of the array are created.


*/
type MultiANewArray struct {
	index      uint16
	dimensions uint8
}

func (its *MultiANewArray) FetchOperand(reader *base.BytecodeReader) {
	its.index = reader.ReadUint16()
	its.dimensions = reader.ReadUint8()
}

func (its *MultiANewArray) Execute(frame *rtdaBase.JavaFrame) {
	if its.dimensions < 1 {
		panic(fmt.Sprintf("multianewarray dimensions must >= 1. (%v)", its.dimensions))
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(its.index)).(*java.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(its.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *rtdaBase.JavaOperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

func newMultiDimensionalArray(counts []int32, arrClass *java.JavaClass) *java.JavaObject {
	count := uint(counts[0])
	arr := arrClass.NewJavaArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
