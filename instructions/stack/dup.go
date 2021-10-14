package stack

import (
	"encoding/json"

	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

/*
Duplicate the top operand stack value
*/
type Dup struct {
}

func (its *Dup) FetchOperand(reader *base.BytecodeReader) {

}

func (its *Dup) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	msg, err := json.Marshal(slot)
	if err != nil {
		mylog.Printf("Slot: %v", *slot)
	} else {
		mylog.Printf("Slot: %v", msg)
	}
}
