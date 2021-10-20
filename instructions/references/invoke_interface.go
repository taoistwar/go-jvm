package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

// Invoke interface method
type InvokeInterface struct {
	index uint
	// 第3字节的值是给方法传递参数需要的slot数，其含义和给Method结构体定义的argSlotCount字段相同。
	// 正如我们所知，这个数是可以根据方法描述符计算出来的，它的存在仅仅是因为历史原因。
	// count uint8
	// 第4字节是留给Oracle的某些Java虚拟机实现用的，它的值必须是0。该字节的存在是为了保证Java虚拟机可以向后兼容。
	// zero  uint8
}

func (its *InvokeInterface) FetchOperand(reader *base.BytecodeReader) {
	its.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (its *InvokeInterface) Execute(frame *rtdaBase.JavaFrame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(its.index).(*java.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	mylog.Printf("Index: %v, Class: %v, Method: %v %v", its.index, resolvedMethod.Class().ThisClassName(), resolvedMethod.Name(), resolvedMethod.Descriptor())

	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中弹出this引用，如果引用是null，则抛出NullPointerException异常。
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException") // todo
	}
	// 如果引用所指对象的类没有实现解析出来的接口，则抛出IncompatibleClassChangeError异常。
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := ref.Class().LookupMethodInClass(methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
