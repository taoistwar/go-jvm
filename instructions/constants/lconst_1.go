package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LConst1 struct {
}

func (its *LConst1) FetchOperand(reader *base.BytecodeReader) {

}
func (its *LConst1) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushLong(1)
}
