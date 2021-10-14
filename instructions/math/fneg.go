package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FNeg struct {
}

func (its *FNeg) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FNeg) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	stack.PushFloat(-v1)
}
