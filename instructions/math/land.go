package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LAnd struct {
}

func (its *LAnd) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LAnd) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	result := v1 & v2
	stack.PushLong(result)
}
