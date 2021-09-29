package classfile

import "fmt"

// https://docs.oracle.com/javase/specs/jvms/se17/html/jvms-4.html#jvms-4.4.5
// Constant pool tags
const (
	CONSTANT_Utf8               = 1  // 45.3 -> 1.0.2
	CONSTANT_Integer            = 3  // 45.3 -> 1.0.2
	CONSTANT_Float              = 4  // 45.3 -> 1.0.2
	CONSTANT_Long               = 5  // 45.3 -> 1.0.2
	CONSTANT_Double             = 6  // 45.3 -> 1.0.2
	CONSTANT_Class              = 7  // 45.3 -> 1.0.2
	CONSTANT_String             = 8  // 45.3 -> 1.0.2
	CONSTANT_Fieldref           = 9  // 45.3 -> 1.0.2
	CONSTANT_Methodref          = 10 // 45.3 -> 1.0.2
	CONSTANT_InterfaceMethodref = 11 // 45.3 -> 1.0.2
	CONSTANT_NameAndType        = 12 // 45.3 -> 1.0.2
	CONSTANT_MethodHandle       = 15 // 51.0 -> 7
	CONSTANT_MethodType         = 16 // 51.0 -> 7
	CONSTANT_Dynamic            = 17 // 55.0-> 11
	CONSTANT_InvokeDynamic      = 18
	CONSTANT_Module             = 19 // 53.0 -> 9
	CONSTANT_Package            = 20 // 53.0 -> 9
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.ReadUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_Dynamic:
		return &ConstantDynamicInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	case CONSTANT_Module:
		return &ConstantModuleInfo{}
	case CONSTANT_Package:
		return &ConstantPackageInfo{}
	default:
		msg := fmt.Sprintf("java.lang.ClassFormatError: constant pool tag(%v)!", tag)
		panic(msg)
	}
}
