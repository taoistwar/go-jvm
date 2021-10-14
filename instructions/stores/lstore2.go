package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LStore2 struct {
}

func (its *LStore2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LStore2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(2, value)
}
