package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LStore3 struct {
}

func (its *LStore3) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LStore3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(3, value)
}
