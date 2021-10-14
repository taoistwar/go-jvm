package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FStore3 struct {
}

func (its *FStore3) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FStore3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(3, value)
}
