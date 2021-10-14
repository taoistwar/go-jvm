package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad1 struct {
}

func (its *ALoad1) FetchOperand(reader *base.BytecodeReader) {

}

func (its *ALoad1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(1)
	frame.OperandStack().PushRef(value)
}
