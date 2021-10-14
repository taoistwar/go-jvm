package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IConst2 struct {
}

func (its *IConst2) FetchOperand(reader *base.BytecodeReader) {

}
func (its *IConst2) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushInt(2)
}
