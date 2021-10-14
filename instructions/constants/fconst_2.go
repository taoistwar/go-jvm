package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FConst2 struct {
}

func (its *FConst2) FetchOperand(reader *base.BytecodeReader) {

}
func (its *FConst2) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushFloat(2)
}
