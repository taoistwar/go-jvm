package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type D2I struct {
}

func (its *D2I) FetchOperand(reader *base.BytecodeReader) {

}

func (its *D2I) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}
