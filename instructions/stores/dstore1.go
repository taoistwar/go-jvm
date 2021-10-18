package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DStore1 struct {
}

func (its *DStore1) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DStore1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(1, value)
}
