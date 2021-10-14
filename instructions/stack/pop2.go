package stack

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type Pop2 struct {
}

func (its *Pop2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *Pop2) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
