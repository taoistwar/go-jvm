package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad3 struct {
}

func (its *ALoad3) FetchOperand(reader *base.BytecodeReader) {

}

func (its *ALoad3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(3)
	frame.OperandStack().PushRef(value)
}
