package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IConst4 struct {
}

func (its *IConst4) FetchOperand(reader *base.BytecodeReader) {

}
func (its *IConst4) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushInt(4)
}
