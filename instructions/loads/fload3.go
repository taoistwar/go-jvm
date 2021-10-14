package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FLoad3 struct {
}

func (its *FLoad3) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FLoad3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetFloat(3)
	frame.OperandStack().PushFloat(value)
}
