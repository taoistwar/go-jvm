package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DDiv struct {
}

func (its *DDiv) FetchOperand(reader *base.BytecodeReader) {

}

func (its *DDiv) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	if v2 == 0 {
		panic("java.lang.ArithmeticExcpetion: / by zero")
	}
	result := v1 / v2
	stack.PushDouble(result)
}
