package java

import "github.com/taoistwar/go-jvm/classfile"

type JavaField struct {
	accessFlags     uint16
	name            string
	descriptor      string
	class           *JavaClass
	constValueIndex uint
	slotId          uint
}

func newFields(class *JavaClass, cfFields []*classfile.MemberInfo) []*JavaField {
	fields := make([]*JavaField, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &JavaField{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (its *JavaField) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	its.accessFlags = memberInfo.AccessFlags()
	its.name = memberInfo.Name()
	its.descriptor = memberInfo.Descriptor()
}

func (its *JavaField) IsPublic() bool {
	return its.accessFlags&ACC_PUBLIC != 0
}
func (its *JavaField) IsPrivate() bool {
	return its.accessFlags&ACC_PRIVATE != 0
}
func (its *JavaField) IsProtected() bool {
	return its.accessFlags&ACC_PROTECTED != 0
}
func (its *JavaField) IsStatic() bool {
	return its.accessFlags&ACC_STATIC != 0
}
func (its *JavaField) IsFinal() bool {
	return its.accessFlags&ACC_FINAL != 0
}
func (its *JavaField) IsSynthetic() bool {
	return its.accessFlags&ACC_SYNTHETIC != 0
}

// getters
func (its *JavaField) Name() string {
	return its.name
}
func (its *JavaField) Descriptor() string {
	return its.descriptor
}
func (its *JavaField) Class() *JavaClass {
	return its.class
}

// jvms 5.4.4
func (its *JavaField) isAccessibleTo(d *JavaClass) bool {
	if its.IsPublic() {
		return true
	}
	c := its.class
	if its.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !its.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (its *JavaField) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		its.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (its *JavaField) IsVolatile() bool {
	return its.accessFlags&ACC_VOLATILE != 0
}
func (its *JavaField) IsTransient() bool {
	return its.accessFlags&ACC_TRANSIENT != 0
}
func (its *JavaField) IsEnum() bool {
	return its.accessFlags&ACC_ENUM != 0
}

func (its *JavaField) ConstValueIndex() uint {
	return its.constValueIndex
}
func (its *JavaField) SlotId() uint {
	return its.slotId
}
func (its *JavaField) isLongOrDouble() bool {
	return its.descriptor == "J" || its.descriptor == "D"
}
