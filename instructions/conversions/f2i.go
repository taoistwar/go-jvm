package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type F2I struct {
}

func (its *F2I) FetchOperand(reader *base.BytecodeReader) {

}

func (its *F2I) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
