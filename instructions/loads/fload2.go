package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FLoad2 struct {
}

func (its *FLoad2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FLoad2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetFloat(2)
	frame.OperandStack().PushFloat(value)
}
