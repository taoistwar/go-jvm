package control

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FReturn struct {
}

func (its *FReturn) FetchOperand(reader *base.BytecodeReader) {
}

func (its *FReturn) Execute(frame *rtdaBase.JavaFrame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}
