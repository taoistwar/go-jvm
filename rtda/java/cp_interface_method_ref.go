package java

import "github.com/taoistwar/go-jvm/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *JavaMethod
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (imr *InterfaceMethodRef) ResolvedInterfaceMethod() *JavaMethod {
	if imr.method == nil {
		imr.resolveInterfaceMethodRef()
	}
	return imr.method
}

// jvms8 5.4.3.4
func (its *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := its.cp.class
	c := its.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := c.LookupInterfaceMethod(its.name, its.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	its.method = method
}
