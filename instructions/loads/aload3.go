package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ALoad3 struct {
}

func (its *ALoad3) FetchOperand(reader *base.BytecodeReader) {

}

func (its *ALoad3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetRef(3)
	frame.OperandStack().PushRef(value)
	mylog.Printf("Index: 3, Class: %v, Fields: %v, Value: %v", value.Class(), value.Fields(), value)
}
