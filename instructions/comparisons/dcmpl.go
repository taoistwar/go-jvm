package comparisons

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DCmpL struct {
}

func (its *DCmpL) FetchOperand(reader *base.BytecodeReader) {

}

func (its *DCmpL) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(-1)
	}
}
