package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FLoad1 struct {
}

func (its *FLoad1) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FLoad1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetFloat(1)
	frame.OperandStack().PushFloat(value)
}
