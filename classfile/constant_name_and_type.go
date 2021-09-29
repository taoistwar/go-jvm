package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (its *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	its.nameIndex = reader.ReadUint16()
	its.descriptorIndex = reader.ReadUint16()
}
