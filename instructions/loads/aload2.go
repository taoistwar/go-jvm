package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad2 struct {
}

func (its *ALoad2) FetchOperand(reader *base.BytecodeReader) {

}

func (its *ALoad2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(2)
	mylog.Printf("Value: %v", value)
	frame.OperandStack().PushRef(value)
}
