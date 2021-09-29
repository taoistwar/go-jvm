package classfile

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (its *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint32()
	its.val = int32(bytes)
}
func (its *ConstantIntegerInfo) Value() int32 {
	return its.val
}
