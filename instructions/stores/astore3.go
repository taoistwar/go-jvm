package stores

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type AStore3 struct {
	index int8
}

func (its *AStore3) FetchOperand(reader *base.BytecodeReader) {
	its.index = reader.ReadInt8()
}

func (its *AStore3) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.OperandStack().PopRef()
	mylog.Printf("Index:%v, Value: %v", its.index, value)
	frame.LocalVars().SetRef(3, value)
}
