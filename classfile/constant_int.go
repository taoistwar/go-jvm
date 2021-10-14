package classfile

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	value int32
}

func (its *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	// The bytes item of the CONSTANT_Integer_info structure represents the value of the int constant.
	// The bytes of the value are stored in big-endian (high byte first) order.
	bytes := reader.ReadUint32()
	its.value = int32(bytes)
}
func (its *ConstantIntegerInfo) Value() int32 {
	return its.value
}
