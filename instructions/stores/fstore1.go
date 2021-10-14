package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FStore1 struct {
}

func (its *FStore1) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FStore1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(1, value)
}
