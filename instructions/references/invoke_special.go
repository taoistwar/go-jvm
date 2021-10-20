package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type InvokeSpecial struct {
	Index uint
}

func (its *InvokeSpecial) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadUint16())
}

func (its *InvokeSpecial) Execute(frame *rtdaBase.JavaFrame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(its.Index).(*java.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	mylog.Printf("Index: %v, Class: %v, Method: %v %v", its.Index, resolvedClass.ThisClassName(), resolvedMethod.Name(), resolvedMethod.Descriptor())

	// 假定从方法符号引用中解析出来的类是C，方法是M。
	// 如果M是构造函数，则声明M的类必须是C，否则抛出NoSuchMethodError异常。
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	// 如果M是静态方法，则抛出IncompatibleClassChangeError异常。
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中弹出this引用，如果该引用是null，抛出NullPointerException异常。
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// protected方法只能被声明该方法的类或子类调用。
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	// 如果调用的超类中的函数，但不是构造函数，且当前类的ACC_SUPER标志被设置，需要一个额外的过程查找最终要调用的方法；
	// 否则前面从方法符号引用中解析出来的方法就是要调用的方法。
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = currentClass.SuperClass().LookupMethodInClass(methodRef.Name(), methodRef.Descriptor())
	}

	// 如果查找过程失败，或者找到的方法是抽象的，抛出AbstractMethodError异常。
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	// 这里之所以这么复杂，是因为调用超类的（非构造函数）方法需要特别处理。
	// 限于篇幅，这里就不深入讨论了，读者可以阅读Java虚拟机规范，了解类的 ACC_SUPER 访问标志的用法。
	base.InvokeMethod(frame, methodToBeInvoked)
}
