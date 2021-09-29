package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (its *ConstantStringInfo) readInfo(reader *ClassReader) {
	its.stringIndex = reader.ReadUint16()
}
func (its *ConstantStringInfo) String() string {
	return its.cp.getUtf8(its.stringIndex)
}
