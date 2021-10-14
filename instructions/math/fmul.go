package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FMul struct {
}

func (its *FMul) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FMul) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}
