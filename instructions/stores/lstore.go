package stores

import (
	instBase "github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LStore struct {
	Index uint
}

func (its *LStore) FetchOperand(reader *instBase.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *LStore) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(uint(its.Index), value)
}
