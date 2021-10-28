package java

import (
	"fmt"
	"strings"

	"github.com/taoistwar/go-jvm/classfile"
)

type JavaClass struct {
	accessFlags       uint16
	thisClassName     string        // 完全限定名，具有java/lang/Object的形式。
	superClassName    string        // 完全限定名，具有java/lang/Object的形式。
	interfaceNames    []string      // 完全限定名，具有java/lang/Object的形式。
	constantPool      *ConstantPool // 运行时常量池指针
	fields            []*JavaField  // 字段表
	methods           []*JavaMethod // 方法表
	loader            *JavaClassLoader
	superClass        *JavaClass
	interfaces        []*JavaClass
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *JavaObject
}

func NewJavaClass(cf *classfile.ClassFile) *JavaClass {
	class := &JavaClass{}
	class.accessFlags = cf.AccessFlags()
	class.thisClassName = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}
func (its *JavaClass) IsArray() bool {
	return its.thisClassName[0] == '['
}
func (its *JavaClass) NewJavaArray(count uint) *JavaObject {
	if !its.IsArray() {
		panic(fmt.Sprintf("Not array class: %v", its.thisClassName))
	}
	switch its.thisClassName {
	case "[Z": // boolean
		return &JavaObject{class: its, data: make([]int8, count)}
	case "[B": // byte
		return &JavaObject{class: its, data: make([]int8, count)}
	case "[C": // chart
		return &JavaObject{class: its, data: make([]uint16, count)}
	case "[S": // short
		return &JavaObject{class: its, data: make([]int16, count)}
	case "[I": // int
		return &JavaObject{class: its, data: make([]int32, count)}
	case "[J": // long
		return &JavaObject{class: its, data: make([]int64, count)}
	case "[F": // float
		return &JavaObject{class: its, data: make([]float32, count)}
	case "[D": // double
		return &JavaObject{class: its, data: make([]float64, count)}
	default:
		return &JavaObject{class: its, data: make([]*JavaObject, count)}
	}
}

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func toDescriptor(className string) string {
	if className[0] == '[' {
		// array
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		// primitive
		return d
	}
	// object
	return "L" + className + ";"
}
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}
func (its *JavaClass) ArrayClass() *JavaClass {
	arrayClassName := getArrayClassName(its.thisClassName)
	return its.loader.LoadJClass(arrayClassName)
}
func (its *JavaClass) ThisClassName() string {
	return its.thisClassName
}
func (its *JavaClass) IsPublic() bool {
	return its.accessFlags&ACC_PUBLIC != 0
}
func (its *JavaClass) IsFinal() bool {
	return its.accessFlags&ACC_FINAL != 0
}
func (its *JavaClass) IsSuper() bool {
	return its.accessFlags&ACC_SUPER != 0
}
func (its *JavaClass) IsInterface() bool {
	return its.accessFlags&ACC_INTERFACE != 0
}
func (its *JavaClass) IsAbstract() bool {
	return its.accessFlags&ACC_ABSTRACT != 0
}
func (its *JavaClass) IsSynthetic() bool {
	return its.accessFlags&ACC_SYNTHETIC != 0
}
func (its *JavaClass) IsAnnotation() bool {
	return its.accessFlags&ACC_ANNOTATION != 0
}
func (its *JavaClass) IsEnum() bool {
	return its.accessFlags&ACC_ENUM != 0
}

// getters
func (its *JavaClass) ConstantPool() *ConstantPool {
	return its.constantPool
}
func (its *JavaClass) StaticVars() Slots {
	return its.staticVars
}

// jvms 5.4.4
func (its *JavaClass) isAccessibleTo(other *JavaClass) bool {
	return its.IsPublic() ||
		its.getPackageName() == other.getPackageName()
}

func (its *JavaClass) getPackageName() string {
	if i := strings.LastIndex(its.thisClassName, "/"); i >= 0 {
		return its.thisClassName[:i]
	}
	return ""
}

func (its *JavaClass) GetMainJMethod() *JavaMethod {
	return its.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (its *JavaClass) GetClassInitMethod() *JavaMethod {
	return its.getStaticMethod("<clinit>", "()V")
}
func (its *JavaClass) getStaticMethod(name, descriptor string) *JavaMethod {
	for _, method := range its.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (its *JavaClass) NewObject() *JavaObject {
	return newJavaObject(its)
}

// its extends c
func (its *JavaClass) isSubClassOf(other *JavaClass) bool {
	for c := its.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// its implements iface
func (its *JavaClass) isImplements(iface *JavaClass) bool {
	for c := its; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// its extends iface
func (its *JavaClass) isSubInterfaceOf(iface *JavaClass) bool {
	for _, superInterface := range its.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (c *JavaClass) lookupField(name, descriptor string) *JavaField {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := iface.lookupField(name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return c.superClass.lookupField(name, descriptor)
	}
	return nil
}

func (class *JavaClass) LookupMethod(name, descriptor string) *JavaMethod {
	method := class.LookupMethodInClass(name, descriptor)
	if method != nil {
		return method
	}
	for _, iface := range class.interfaces {
		method = iface.LookupMethodInInterface(name, descriptor)
		if method != nil {
			break
		}
	}
	return nil
}

func (class *JavaClass) LookupMethodInClass(name, descriptor string) *JavaMethod {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func (iface *JavaClass) LookupMethodInInterface(name, descriptor string) *JavaMethod {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	for _, iface2 := range iface.interfaces {
		method := iface2.LookupMethodInInterface(name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}

func (its *JavaClass) LookupInterfaceMethod(name, descriptor string) *JavaMethod {
	for _, method := range its.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	for _, iface2 := range its.interfaces {
		method := iface2.LookupInterfaceMethod(name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
func (its *JavaClass) InitializationNotStarted() bool {
	return !its.initStarted
}
func (its *JavaClass) InitStarted() bool {
	return its.initStarted
}
func (its *JavaClass) StartInit() {
	its.initStarted = true
}
func (its *JavaClass) SuperClass() *JavaClass {
	return its.superClass
}
func (its *JavaClass) Interfaces() []*JavaClass {
	return its.interfaces
}
func (its *JavaClass) Methods() []*JavaMethod {
	return its.methods
}
func (its *JavaClass) Fields() []*JavaField {
	return its.fields
}
func (its *JavaClass) Loader() *JavaClassLoader {
	return its.loader
}

// c extends its
func (its *JavaClass) IsSuperClassOf(other *JavaClass) bool {
	return other.IsSubClassOf(its)
}

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
/*
other是否为its的子类
@param other
*/
func (its *JavaClass) IsAssignableFrom(other *JavaClass) bool {
	s, t := other, its

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}
}

// its extends c
func (its *JavaClass) IsSubClassOf(other *JavaClass) bool {
	for c := its.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// its implements iface
func (its *JavaClass) IsImplements(iface *JavaClass) bool {
	for c := its; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}
func (its *JavaClass) GetPackageName() string {
	if i := strings.LastIndex(its.thisClassName, "/"); i >= 0 {
		return its.thisClassName[:i]
	}
	return ""
}
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}
func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		// array
		return descriptor
	}
	if descriptor[0] == 'L' {
		// object
		return descriptor[1 : len(descriptor)-1]
	}
	for className, d := range primitiveTypes {
		if d == descriptor {
			// primitive
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
func (its *JavaClass) ComponentClass() *JavaClass {
	componentClassName := getComponentClassName(its.thisClassName)
	return its.loader.LoadJClass(componentClassName)
}

func (its *JavaClass) getField(name, descriptor string, isStatic bool) *JavaField {
	for c := its; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (its *JavaClass) isJlObject() bool {
	return its.thisClassName == "java/lang/Object"
}
func (its *JavaClass) isJlCloneable() bool {
	return its.thisClassName == "java/lang/Cloneable"
}
func (its *JavaClass) isJioSerializable() bool {
	return its.thisClassName == "java/io/Serializable"
}
func (its *JavaClass) JClass() *JavaObject {
	return its.jClass
}

// iface extends self
func (its *JavaClass) isSuperInterfaceOf(iface *JavaClass) bool {
	return iface.isSubInterfaceOf(its)
}
func (its *JavaClass) GetRefVar(fieldName, fieldDescriptor string) *JavaObject {
	field := its.getField(fieldName, fieldDescriptor, true)
	return its.staticVars.GetRef(field.slotId)
}
func (its *JavaClass) SetRefVar(fieldName, fieldDescriptor string, ref *JavaObject) {
	field := its.getField(fieldName, fieldDescriptor, true)
	its.staticVars.SetRef(field.slotId, ref)
}
