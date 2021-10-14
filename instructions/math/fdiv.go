package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FDiv struct {
}

func (its *FDiv) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FDiv) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if v2 == 0 {
		panic("java.lang.ArithmeticExcpetion: / by zero")
	}
	result := v1 / v2
	stack.PushFloat(result)
}
