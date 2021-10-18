package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AStore1 struct {
}

func (its *AStore1) FetchOperand(reader *base.BytecodeReader) {
}

func (its *AStore1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(1, value)
}
