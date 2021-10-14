package constants

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IConstM1 struct {
}

func (its *IConstM1) FetchOperand(reader *base.BytecodeReader) {

}
func (its *IConstM1) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PushInt(-1)
}
