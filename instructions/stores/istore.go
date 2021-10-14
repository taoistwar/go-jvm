package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IStore struct {
	Index uint
}

func (its *IStore) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *IStore) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(uint(its.Index), value)
}
