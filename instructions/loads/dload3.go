package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DLoad3 struct {
}

func (its *DLoad3) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DLoad3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetDouble(3)
	frame.OperandStack().PushDouble(value)
}
