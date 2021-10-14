package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DNeg struct {
}

func (its *DNeg) FetchOperand(reader *base.BytecodeReader) {

}

func (its *DNeg) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	stack.PushDouble(-v1)
}
