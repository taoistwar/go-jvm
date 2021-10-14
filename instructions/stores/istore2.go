package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IStore2 struct {
}

func (its *IStore2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IStore2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(2, value)
}
