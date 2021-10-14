package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DStore struct {
	Index uint
}

func (its *DStore) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *DStore) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(uint(its.Index), value)
}
