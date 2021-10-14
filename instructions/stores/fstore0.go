package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FStore0 struct {
}

func (its *FStore0) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FStore0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(0, value)
}
