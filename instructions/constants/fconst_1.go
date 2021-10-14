package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FConst1 struct {
}

func (its *FConst1) FetchOperand(reader *base.BytecodeReader) {

}
func (its *FConst1) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushFloat(1)
}
