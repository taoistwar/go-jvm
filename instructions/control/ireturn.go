package control

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IReturn struct {
}

func (its *IReturn) FetchOperand(reader *base.BytecodeReader) {
}

func (its *IReturn) Execute(frame *rtdaBase.JavaFrame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}
