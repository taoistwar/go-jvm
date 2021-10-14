package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IRem struct {
}

func (its *IRem) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IRem) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticExcpetion: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}
