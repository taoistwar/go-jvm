package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IConst0 struct {
}

func (its *IConst0) FetchOperand(reader *base.BytecodeReader) {

}
func (its *IConst0) Execute(frame *rtdaBase.JavaFrame) {

	frame.OperandStack().PushInt(0)
}
