package java

import "github.com/taoistwar/go-jvm/classfile"

type MethodRef struct {
	MemberRef
	method *JavaMethod
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (its *MethodRef) ResolvedMethod() *JavaMethod {
	if its.method == nil {
		its.resolveMethodRef()
	}
	return its.method
}

// jvms8 5.4.3.3
func (its *MethodRef) resolveMethodRef() {
	d := its.cp.class
	c := its.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := c.LookupMethod(its.name, its.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	its.method = method
}
