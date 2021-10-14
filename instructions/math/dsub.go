package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DSub struct {
}

func (its *DSub) FetchOperand(reader *base.BytecodeReader) {

}

func (its *DSub) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}
