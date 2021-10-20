package java

import (
	"fmt"

	"github.com/taoistwar/go-jvm/classfile"
)

type ConstantPool struct {
	class     *JavaClass
	constants []Constant
}

func newConstantPool(class *JavaClass, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	constants := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, constants}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo := cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			constants[i] = cpInfo.Value()
		case *classfile.ConstantFloatInfo:
			constants[i] = cpInfo.Value()
		case *classfile.ConstantLongInfo:
			constants[i] = cpInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			constants[i] = cpInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			constants[i] = cpInfo.String()
		case *classfile.ConstantClassInfo:
			constants[i] = newClassRef(rtCp, cpInfo)
		case *classfile.ConstantFieldrefInfo:
			constants[i] = newFieldRef(rtCp, cpInfo)
		case *classfile.ConstantMethodrefInfo:
			constants[i] = newMethodRef(rtCp, cpInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			constants[i] = newInterfaceMethodRef(rtCp, cpInfo)
		default:
			// todo
		}
	}

	return rtCp
}

func (cp *ConstantPool) GetConstant(index uint) Constant {
	if c := cp.constants[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
