package base

type BytecodeReader struct {
	code []byte // bytecodes
	pc   int
}

func (its *BytecodeReader) Reset(code []byte, pc int) {
	its.code = code
	its.pc = pc
}

func (its *BytecodeReader) PC() int {
	return its.pc
}

func (its *BytecodeReader) ReadInt8() int8 {
	return int8(its.ReadOperandCode())
}
func (its *BytecodeReader) ReadOperandCode() uint8 {
	i := its.code[its.pc]
	its.pc++
	return i
}
func (its *BytecodeReader) ReadUint8() uint8 {
	i := its.code[its.pc]
	its.pc++
	return i
}

func (its *BytecodeReader) ReadInt16() int16 {
	return int16(its.ReadUint16())
}
func (its *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(its.ReadOperandCode())
	byte2 := uint16(its.ReadOperandCode())
	return (byte1 << 8) | byte2
}

func (its *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(its.ReadOperandCode())
	byte2 := int32(its.ReadOperandCode())
	byte3 := int32(its.ReadOperandCode())
	byte4 := int32(its.ReadOperandCode())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// used by lookupswitch and tableswitch
func (its *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = its.ReadInt32()
	}
	return ints
}

// used by lookupswitch and tableswitch
func (its *BytecodeReader) SkipPadding() {
	for its.pc%4 != 0 {
		its.ReadOperandCode()
	}
}
