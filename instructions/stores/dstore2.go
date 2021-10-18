package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DStore2 struct {
}

func (its *DStore2) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DStore2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(2, value)
}
