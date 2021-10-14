package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad0 struct {
}

func (its *ALoad0) FetchOperand(reader *base.BytecodeReader) {

}

func (its *ALoad0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(0)
	frame.OperandStack().PushRef(value)
}
