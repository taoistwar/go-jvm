package control

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type BytecodeReturn struct {
}

func (its *BytecodeReturn) FetchOperand(reader *base.BytecodeReader) {
}

func (its *BytecodeReturn) Execute(frame *rtdaBase.JavaFrame) {
	frame.Thread().PopFrame()
}
