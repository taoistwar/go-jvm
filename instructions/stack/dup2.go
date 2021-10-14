package stack

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type Dup2 struct {
}

func (its *Dup2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *Dup2) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
