package control

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AReturn struct {
}

func (its *AReturn) FetchOperand(reader *base.BytecodeReader) {
}

func (its *AReturn) Execute(frame *rtdaBase.JavaFrame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(val)
}
