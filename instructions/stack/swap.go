package stack

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type Swap struct {
}

func (its *Swap) FetchOperand(reader *base.BytecodeReader) {

}

func (its *Swap) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
