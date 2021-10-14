package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LNeg struct {
}

func (its *LNeg) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LNeg) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	stack.PushLong(-v1)
}
