package classfile

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (its *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	its.val = int64(bytes)
}
func (its *ConstantLongInfo) Value() int64 {
	return its.val
}
