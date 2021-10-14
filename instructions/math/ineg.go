package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type INeg struct {
}

func (its *INeg) FetchOperand(reader *base.BytecodeReader) {

}

func (its *INeg) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	stack.PushInt(-v1)
}
