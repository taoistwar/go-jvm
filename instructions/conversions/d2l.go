package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type D2L struct {
}

func (its *D2L) FetchOperand(reader *base.BytecodeReader) {

}

func (its *D2L) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
