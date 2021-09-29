package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (its *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	its.classIndex = reader.ReadUint16()
	its.nameAndTypeIndex = reader.ReadUint16()
}

func (its *ConstantMemberrefInfo) ClassName() string {
	return its.cp.getClassName(its.classIndex)
}
func (its *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return its.cp.getNameAndType(its.nameAndTypeIndex)
}
