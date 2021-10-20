package java

import (
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

func (its *JavaClass) GetMainMethod() *JavaMethod {
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

func (its *JavaClass) InitStarted() bool {
	return its.initStarted
}
func (its *JavaClass) StartInit() {
	its.initStarted = true
}
func (its *JavaClass) SuperClass() *JavaClass {
	return its.superClass
}

// c extends self
func (its *JavaClass) IsSuperClassOf(other *JavaClass) bool {
	return other.IsSubClassOf(its)
}

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (its *JavaClass) isAssignableFrom(other *JavaClass) bool {
	s, t := other, its

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

// self extends c
func (its *JavaClass) IsSubClassOf(other *JavaClass) bool {
	for c := its.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
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
