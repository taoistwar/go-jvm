package classfile

import (
	"fmt"
)

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	constantPoolCount := reader.ReadUint16()
	cp := make([]ConstantInfo, constantPoolCount)

	// The constant_pool table is indexed from 1 to constant_pool_count - 1.
	var i uint16 = 1
	for ; i < constantPoolCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		// https://docs.oracle.com/javase/specs/jvms/se17/html/jvms-4.html#jvms-4.4.5
		// All 8-byte constants take up two entries in the constant_pool table of the class file.
		// If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
		// table at index n, then the next usable item in the pool is located at index n+2.
		// The constant_pool index n+1 must be valid but is considered unusable.
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

func (its ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := its[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Sprintf("Invalid constant pool index: %v!", index))
}

func (its ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := its.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := its.getUtf8(ntInfo.nameIndex)
	_type := its.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (its ConstantPool) getClassName(index uint16) string {
	cpInfo := its.getConstantInfo(index)
	classInfo := cpInfo.(*ConstantClassInfo)
	return its.getUtf8(classInfo.nameIndex)
}

func (its ConstantPool) getUtf8(index uint16) string {
	cpInfo := its.getConstantInfo(index)
	utf8Info := cpInfo.(*ConstantUtf8Info)
	return utf8Info.str
}
