package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (its *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	its.referenceKind = reader.ReadUint8()
	its.referenceIndex = reader.ReadUint16()
}
