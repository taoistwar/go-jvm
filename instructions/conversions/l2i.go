package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type L2I struct {
}

func (its *L2I) FetchOperand(reader *base.BytecodeReader) {

}

func (its *L2I) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
