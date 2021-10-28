package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AStore3 struct {
}

func (its *AStore3) FetchOperand(reader *base.BytecodeReader) {
}

func (its *AStore3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopRef()
	mylog.Printf("Index: 3, Class: %v, Value: %v", value.Class().ThisClassName(), value)
	frame.LocalVars().SetRef(3, value)
}
