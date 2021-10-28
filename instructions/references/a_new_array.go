package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

/*
Operation
	Create new array of reference

Format
	anewarray indexbyte1 indexbyte2

Forms
	anewarray = 189 (0xbd)

Operand Stack
	..., count →
	..., arrayref

Description
	The count must be of type int. 计数必须为int类型。
	It is popped off the operand stack. 它从操作数堆栈中弹出。
	The count represents the number of components of the array to be created.该计数表示要创建的数组的元素数量。
	The unsigned indexbyte1 and indexbyte2 are used to construct an index into the run-time constant pool of the current class (§2.6), where the value of the index is (indexbyte1 << 8) | indexbyte2.
	无符号indexbyte1和indexbyte2用于构造当前类的运行时常量池的索引(§2.6)，其中索引的值是(indexbyte1 << 8) | indexbyte2。
	The run-time constant pool item at that index must be a symbolic reference to a class, array, or interface type.
	该索引处的运行时常量池项必须是对类、数组或接口类型的符号引用。
	The named class, array, or interface type is resolved (§5.4.3.1).解析命名的类、数组或接口类型(§5.4.3.1)。
	A new array with components of that type, of length count, is allocated from the garbage-collected heap, and a reference arrayref to this new array object is pushed onto the operand stack.
	从垃圾收集堆中分配一个包含该类型组件(长度计数)的新数组，并将此新数组对象的引用arrayref推入操作数堆栈。
	All components of the new array are initialized to null, the default value for reference types (§2.4).
	新数组的所有组件都被初始化为null，这是引用类型的默认值(§2.4)。

Linking Exceptions
	During resolution of the symbolic reference to the class, array, or interface type, any of the exceptions documented in §5.4.3.1 can be thrown.
	在解析类、数组或接口类型的符号引用时，可以抛出§5.4.3.1中记录的任何异常。

Run-time Exceptions
	Otherwise, if count is less than zero, the anewarray instruction throws a NegativeArraySizeException.
	否则，如果count小于0,anewarray指令抛出一个NegativeArraySizeException。

Notes
	The anewarray instruction is used to create a single dimension of an array of object references or part of a multidimensional array.
	anewarray指令用于创建对象引用数组的一维或多维数组的一部分。
*/
type ANewArray struct {
	Index uint
}

func (its *ANewArray) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadUint16())
}

func (its *ANewArray) Execute(frame *rtdaBase.JavaFrame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(its.Index).(*java.ClassRef)
	componentClass := classRef.ResolvedClass()

	// if componentClass.InitializationNotStarted() {
	// 	thread := frame.Thread()
	// 	frame.SetNextPC(thread.PC()) // undo anewarray
	// 	thread.InitClass(componentClass)
	// 	return
	// }

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewJavaArray(uint(count))
	stack.PushRef(arr)
}
