package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DStore0 struct {
}

func (its *DStore0) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DStore0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(0, value)
}
