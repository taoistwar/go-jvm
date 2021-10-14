package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LLoad1 struct {
}

func (its *LLoad1) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LLoad1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(1)
	frame.OperandStack().PushInt(value)
}
