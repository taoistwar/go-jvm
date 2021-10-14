package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AStore2 struct {
	index int8
}

func (its *AStore2) FetchOperand(reader *base.BytecodeReader) {
	its.index = reader.ReadInt8()
}

func (its *AStore2) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(2, value)
	mylog.Printf("Index: %v, Value: %v", its.index, value.Class().ThisClassName())
}
