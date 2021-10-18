package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AStore0 struct {
}

func (its *AStore0) FetchOperand(reader *base.BytecodeReader) {
}

func (its *AStore0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(0, value)
}
