package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LLoad3 struct {
}

func (its *LLoad3) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LLoad3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(3)
	frame.OperandStack().PushInt(value)
}
