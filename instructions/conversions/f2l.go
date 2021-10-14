package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type F2L struct {
}

func (its *F2L) FetchOperand(reader *base.BytecodeReader) {

}

func (its *F2L) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
