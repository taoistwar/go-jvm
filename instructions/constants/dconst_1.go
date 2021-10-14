package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DConst1 struct {
}

func (its *DConst1) FetchOperand(reader *base.BytecodeReader) {

}
func (its *DConst1) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushDouble(1)
}
