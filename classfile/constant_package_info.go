package classfile

/*
CONSTANT_Package_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantPackageInfo struct {
	nameIndex uint16
}

func (its *ConstantPackageInfo) readInfo(reader *ClassReader) {
	its.nameIndex = reader.ReadUint16()
}
