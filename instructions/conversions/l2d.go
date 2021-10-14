package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type L2D struct {
}

func (its *L2D) FetchOperand(reader *base.BytecodeReader) {

}

func (its *L2D) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}
