package conversions

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type F2D struct {
}

func (its *F2D) FetchOperand(reader *base.BytecodeReader) {

}

func (its *F2D) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}
