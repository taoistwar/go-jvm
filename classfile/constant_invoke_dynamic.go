package classfile

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (its *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	its.bootstrapMethodAttrIndex = reader.ReadUint16()
	its.nameAndTypeIndex = reader.ReadUint16()
}
