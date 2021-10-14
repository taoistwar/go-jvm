package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DStore1 struct {
	index int8
}

func (its *DStore1) FetchOperand(reader *base.BytecodeReader) {
	its.index = reader.ReadInt8()
}

func (its *DStore1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(1, value)
}
