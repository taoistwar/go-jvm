package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IStore1 struct {
}

func (its *IStore1) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IStore1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopInt()
	mylog.Printf("Value: %v", value)
	frame.LocalVars().SetInt(1, value)
}
