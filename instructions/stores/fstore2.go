package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FStore2 struct {
}

func (its *FStore2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FStore2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(2, value)
}
