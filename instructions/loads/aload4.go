package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad4 struct {
}

func (its *ALoad4) FetchOperand(reader *base.BytecodeReader) {

}

func (its *ALoad4) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(4)
	frame.OperandStack().PushRef(value)
}
