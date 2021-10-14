package math

import (
	"math"

	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FRem struct {
}

func (its *FRem) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FRem) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if v2 == 0 {
		panic("java.lang.ArithmeticExcpetion: / by zero")
	}
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}
