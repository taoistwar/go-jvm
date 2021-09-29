package classfile

/*
CONSTANT_Dynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (its *ConstantDynamicInfo) readInfo(reader *ClassReader) {
	its.bootstrapMethodAttrIndex = reader.ReadUint16()
	its.nameAndTypeIndex = reader.ReadUint16()
}
