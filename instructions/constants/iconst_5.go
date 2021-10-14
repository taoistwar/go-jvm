package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IConst5 struct {
}

func (its *IConst5) FetchOperand(reader *base.BytecodeReader) {

}
func (its *IConst5) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushInt(5)
}
