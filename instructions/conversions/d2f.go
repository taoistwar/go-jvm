package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type D2F struct {
}

func (its *D2F) FetchOperand(reader *base.BytecodeReader) {

}

func (its *D2F) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}
