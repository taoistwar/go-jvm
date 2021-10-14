package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type I2L struct {
}

func (its *I2L) FetchOperand(reader *base.BytecodeReader) {

}

func (its *I2L) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}
