package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type L2F struct {
}

func (its *L2F) FetchOperand(reader *base.BytecodeReader) {

}

func (its *L2F) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}
