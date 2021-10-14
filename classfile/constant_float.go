package classfile

import (
	"math"
)

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
The bytes item of the CONSTANT_Float_info structure represents the value of the float constant in IEEE 754 binary32 floating-point format (§2.3.2).
The bytes of the item are stored in big-endian (high byte first) order.
*/
type ConstantFloatInfo struct {
	value float32
}

func (its *ConstantFloatInfo) readInfo(reader *ClassReader) {
	// The bytes item of the CONSTANT_Float_info structure represents the value of the float constant in IEEE 754 binary32 floating-point format
	// (§2.3.2 -> https://docs.oracle.com/javase/specs/jvms/se17/html/jvms-2.html#jvms-2.3.2).
	// The bytes of the item are stored in big-endian (high byte first) order.
	bytes := reader.ReadUint32()
	its.value = math.Float32frombits(bytes)
}
func (its *ConstantFloatInfo) Value() float32 {
	return its.value
}
