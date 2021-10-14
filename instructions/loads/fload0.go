package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FLoad0 struct {
}

func (its *FLoad0) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FLoad0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetFloat(0)
	frame.OperandStack().PushFloat(value)
}
