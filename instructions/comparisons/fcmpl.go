package comparisons

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FCmpL struct {
}

func (its *FCmpL) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FCmpL) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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
