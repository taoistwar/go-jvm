package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ILoad0 struct {
}

func (its *ILoad0) FetchOperand(reader *base.BytecodeReader) {
}

func (its *ILoad0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(0)
	frame.OperandStack().PushInt(value)
}
