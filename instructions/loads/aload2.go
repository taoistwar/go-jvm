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
	frame.OperandStack().PushRef(value)
	mylog.Printf("Index: 2, Class: %v, Fields: %v, Value: %v", value.Class(), value.Fields(), value)
}
