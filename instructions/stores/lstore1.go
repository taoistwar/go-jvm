package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LStore1 struct {
}

func (its *LStore1) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LStore1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(1, value)
}
