package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DLoad2 struct {
}

func (its *DLoad2) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DLoad2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetDouble(2)
	frame.OperandStack().PushDouble(value)
}
