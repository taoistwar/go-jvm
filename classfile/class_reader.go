package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

// u1
func (cr *ClassReader) ReadUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// u2
func (cr *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// u4
func (cr *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

func (cr *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}
func (cr *ClassReader) ReadUint16Array(count uint16) []uint16 {
	s := make([]uint16, count)
	for i := range s {
		s[i] = cr.ReadUint16()
	}
	return s
}
func (cr *ClassReader) ReadUint16s() []uint16 {
	count := cr.ReadUint16()
	s := make([]uint16, count)
	for i := range s {
		s[i] = cr.ReadUint16()
	}
	return s
}

func (cr *ClassReader) ReadBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}
