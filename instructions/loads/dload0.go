package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DLoad0 struct {
}

func (its *DLoad0) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DLoad0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetDouble(0)
	frame.OperandStack().PushDouble(value)
}
