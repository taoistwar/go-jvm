package java

import "github.com/taoistwar/go-jvm/classfile"

type FieldRef struct {
	MemberRef
	field *JavaField
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (fr *FieldRef) ResolvedField() *JavaField {
	if fr.field == nil {
		fr.resolveFieldRef()
	}
	return fr.field
}

// jvms 5.4.3.2
func (fr *FieldRef) resolveFieldRef() {
	d := fr.cp.class
	c := fr.ResolvedClass()
	field := c.lookupField(fr.name, fr.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	fr.field = field
}
