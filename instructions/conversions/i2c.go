package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type I2C struct {
}

func (its *I2C) FetchOperand(reader *base.BytecodeReader) {

}

func (its *I2C) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}
