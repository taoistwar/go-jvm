package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DStore3 struct {
}

func (its *DStore3) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DStore3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(3, value)
}
