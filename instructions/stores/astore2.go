package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AStore2 struct {
}

func (its *AStore2) FetchOperand(reader *base.BytecodeReader) {
}

func (its *AStore2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(2, value)
	mylog.Printf("Index: %v, Value: %v", 2, value.Class().ThisClassName())
}
