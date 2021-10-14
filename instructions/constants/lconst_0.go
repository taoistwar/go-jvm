package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LConst0 struct {
}

func (its *LConst0) FetchOperand(reader *base.BytecodeReader) {

}
func (its *LConst0) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushLong(0)
}
