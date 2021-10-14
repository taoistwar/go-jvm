package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LStore0 struct {
}

func (its *LStore0) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LStore0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(0, value)
}
