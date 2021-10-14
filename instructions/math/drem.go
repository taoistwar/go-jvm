package math

import (
	"math"

	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DRem struct {
}

func (its *DRem) FetchOperand(reader *base.BytecodeReader) {

}

func (its *DRem) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	if v2 == 0 {
		panic("java.lang.ArithmeticExcpetion: / by zero")
	}
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
