package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad0 struct {
}

func (its *ALoad0) FetchOperand(reader *base.BytecodeReader) {

}

func (its *ALoad0) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(0)
	mylog.Printf("Index: 0, Value: %v", value)
	frame.OperandStack().PushRef(value)

}
