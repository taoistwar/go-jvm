package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IAnd struct {
}

func (its *IAnd) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IAnd) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	result := v1 & v2
	stack.PushInt(result)
}
