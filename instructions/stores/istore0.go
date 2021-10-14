package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IStore0 struct {
}

func (its *IStore0) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IStore0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(0, value)
}
