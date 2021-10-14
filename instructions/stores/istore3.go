package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IStore3 struct {
}

func (its *IStore3) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IStore3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(3, value)
}
