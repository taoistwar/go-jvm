package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (its *ConstantClassInfo) readInfo(reader *ClassReader) {
	its.nameIndex = reader.ReadUint16()
}
func (its *ConstantClassInfo) Name() string {
	return its.cp.getUtf8(its.nameIndex)
}
