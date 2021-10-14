package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LLoad2 struct {
}

func (its *LLoad2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LLoad2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(2)
	frame.OperandStack().PushInt(value)
}
