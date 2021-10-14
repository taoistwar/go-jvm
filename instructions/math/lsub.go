package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LSub struct {
}

func (its *LSub) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LSub) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
