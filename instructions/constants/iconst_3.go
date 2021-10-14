package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IConst3 struct {
}

func (its *IConst3) FetchOperand(reader *base.BytecodeReader) {

}
func (its *IConst3) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushInt(3)
}
