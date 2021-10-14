package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DLoad1 struct {
}

func (its *DLoad1) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DLoad1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetDouble(1)
	frame.OperandStack().PushDouble(value)
}
