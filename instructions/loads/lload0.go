package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LLoad0 struct {
}

func (its *LLoad0) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LLoad0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(0)
	frame.OperandStack().PushInt(value)
}
