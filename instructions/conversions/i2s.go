package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type I2S struct {
}

func (its *I2S) FetchOperand(reader *base.BytecodeReader) {

}

func (its *I2S) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}
