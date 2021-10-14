package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DConst0 struct {
}

func (its *DConst0) FetchOperand(reader *base.BytecodeReader) {

}
func (its *DConst0) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushDouble(0.0)
}
