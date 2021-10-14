package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type I2D struct {
}

func (its *I2D) FetchOperand(reader *base.BytecodeReader) {

}

func (its *I2D) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
