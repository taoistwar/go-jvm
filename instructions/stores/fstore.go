package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FStore struct {
	Index uint
}

func (its *FStore) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *FStore) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(uint(its.Index), value)
}
