package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AStore struct {
	Index uint
}

func (its *AStore) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
}

func (its *AStore) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(uint(its.Index), value)
}
