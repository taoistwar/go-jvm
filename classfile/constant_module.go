package classfile

/*
CONSTANT_Module_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantModuleInfo struct {
	nameIndex uint16
}

func (its *ConstantModuleInfo) readInfo(reader *ClassReader) {
	its.nameIndex = reader.ReadUint16()
}
