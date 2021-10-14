package java

import (
	"fmt"

	"github.com/taoistwar/go-jvm/classfile"
)

type Constant interface{}

type ConstantPool struct {
	class     *JavaClass
	constants []Constant
}

func newConstantPool(class *JavaClass, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	constants := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, constants}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo := cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			constants[i] = cpInfo.Value()
		case *classfile.ConstantFloatInfo:
			constants[i] = cpInfo.Value()
		case *classfile.ConstantLongInfo:
			constants[i] = cpInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			constants[i] = cpInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			constants[i] = cpInfo.String()
		case *classfile.ConstantClassInfo:
			constants[i] = newClassRef(rtCp, cpInfo)
		case *classfile.ConstantFieldrefInfo:
			constants[i] = newFieldRef(rtCp, cpInfo)
		case *classfile.ConstantMethodrefInfo:
			constants[i] = newMethodRef(rtCp, cpInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			constants[i] = newInterfaceMethodRef(rtCp, cpInfo)
		default:
			// todo
		}
	}

	return rtCp
}

func (cp *ConstantPool) GetConstant(index uint) Constant {
	if c := cp.constants[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *JavaClass
}

func (its *SymRef) ClassName() string {
	return its.className
}
func (sr *SymRef) ResolvedClass() *JavaClass {
	if sr.class == nil {
		sr.resolveClassRef()
	}
	return sr.class
}

// jvms8 5.4.3.1
func (sr *SymRef) resolveClassRef() {
	d := sr.cp.class
	c := d.loader.LoadClass(sr.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	sr.class = c
}

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
	field := lookupField(c, fr.name, fr.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	fr.field = field
}

func lookupField(c *JavaClass, name, descriptor string) *JavaField {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (mr *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	mr.className = refInfo.ClassName()
	mr.name, mr.descriptor = refInfo.NameAndDescriptor()
}

func (mr *MemberRef) Name() string {
	return mr.name
}
func (mr *MemberRef) Descriptor() string {
	return mr.descriptor
}

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

func (mr *MethodRef) ResolvedMethod() *JavaMethod {
	if mr.method == nil {
		mr.resolveMethodRef()
	}
	return mr.method
}

// jvms8 5.4.3.3
func (mr *MethodRef) resolveMethodRef() {
	//class := self.Class()
	// todo
}

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
func (imr *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := self.ResolveClass()
	// todo
}
