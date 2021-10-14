package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FConst0 struct {
}

func (its *FConst0) FetchOperand(reader *base.BytecodeReader) {

}
func (its *FConst0) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushFloat(0)
}
