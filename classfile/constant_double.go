package classfile

import (
	"math"
)

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (its *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	its.val = math.Float64frombits(bytes)
}
func (its *ConstantDoubleInfo) Value() float64 {
	return its.val
}
