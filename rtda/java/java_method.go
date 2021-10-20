package java

import (
	"github.com/taoistwar/go-jvm/classfile"
)

type JavaMethod struct {
	accessFlags  uint16
	name         string
	descriptor   string
	class        *JavaClass
	maxStack     uint   // 最大操作数
	maxLocals    uint   // 变量个数：参数+局部
	code         []byte // 方法代码，二进制字节码
	argSlotCount uint   // 参数slot数量
}

func newMethods(class *JavaClass, cfMethods []*classfile.MemberInfo) []*JavaMethod {
	methods := make([]*JavaMethod, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &JavaMethod{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}
func (its *JavaMethod) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(its.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		its.argSlotCount++
		if paramType == "J" || paramType == "D" { // 8字节类型占用2个Slot
			its.argSlotCount++
		}
	}
	if !its.IsStatic() {
		its.argSlotCount++ // `this` reference
	}
}
func (jm *JavaMethod) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	jm.accessFlags = memberInfo.AccessFlags()
	jm.name = memberInfo.Name()
	jm.descriptor = memberInfo.Descriptor()
}

func (jm *JavaMethod) IsPublic() bool {
	return jm.accessFlags&ACC_PUBLIC != 0
}
func (jm *JavaMethod) IsPrivate() bool {
	return jm.accessFlags&ACC_PRIVATE != 0
}
func (jm *JavaMethod) IsProtected() bool {
	return jm.accessFlags&ACC_PROTECTED != 0
}
func (jm *JavaMethod) IsStatic() bool {
	return jm.accessFlags&ACC_STATIC != 0
}
func (jm *JavaMethod) IsFinal() bool {
	return jm.accessFlags&ACC_FINAL != 0
}
func (jm *JavaMethod) IsSynthetic() bool {
	return jm.accessFlags&ACC_SYNTHETIC != 0
}

// getters
func (jm *JavaMethod) Name() string {
	return jm.name
}
func (jm *JavaMethod) Descriptor() string {
	return jm.descriptor
}
func (jm *JavaMethod) Class() *JavaClass {
	return jm.class
}

// jvms 5.4.4
func (jm *JavaMethod) isAccessibleTo(d *JavaClass) bool {
	if jm.IsPublic() {
		return true
	}
	c := jm.class
	if jm.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !jm.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (jm *JavaMethod) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		jm.maxStack = codeAttr.MaxStack()
		jm.maxLocals = codeAttr.MaxLocals()
		jm.code = codeAttr.Code()
	}
}

func (jm *JavaMethod) IsSynchronized() bool {
	return jm.accessFlags&ACC_SYNCHRONIZED != 0
}
func (jm *JavaMethod) IsBridge() bool {
	return jm.accessFlags&ACC_BRIDGE != 0
}
func (jm *JavaMethod) IsVarargs() bool {
	return jm.accessFlags&ACC_VARARGS != 0
}
func (jm *JavaMethod) IsNative() bool {
	return jm.accessFlags&ACC_NATIVE != 0
}
func (jm *JavaMethod) IsAbstract() bool {
	return jm.accessFlags&ACC_ABSTRACT != 0
}
func (jm *JavaMethod) IsStrict() bool {
	return jm.accessFlags&ACC_STRICT != 0
}

// getters
func (jm *JavaMethod) MaxStack() uint {
	return jm.maxStack
}
func (jm *JavaMethod) MaxLocals() uint {
	return jm.maxLocals
}
func (jm *JavaMethod) Code() []byte {
	return jm.code
}
func (its *JavaMethod) ArgSlotCount() uint {
	return its.argSlotCount
}
