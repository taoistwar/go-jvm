package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type I2F struct {
}

func (its *I2F) FetchOperand(reader *base.BytecodeReader) {

}

func (its *I2F) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}
