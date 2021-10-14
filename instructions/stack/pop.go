package stack

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type Pop struct {
}

func (its *Pop) FetchOperand(reader *base.BytecodeReader) {

}

func (its *Pop) Execute(frame *rtdaBase.JavaFrame) {
	frame.OperandStack().PopSlot()
}
