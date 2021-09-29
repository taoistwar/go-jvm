package classfile

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (its *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	its.descriptorIndex = reader.ReadUint16()
}
