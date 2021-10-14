package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IConst1 struct {
}

func (its *IConst1) FetchOperand(reader *base.BytecodeReader) {

}
func (its *IConst1) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushInt(1)
}
