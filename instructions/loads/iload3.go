package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ILoad3 struct {
}

func (its *ILoad3) FetchOperand(reader *base.BytecodeReader) {
}

func (its *ILoad3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(3)
	frame.OperandStack().PushInt(value)
}
