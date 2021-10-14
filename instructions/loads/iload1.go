package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type ILoad1 struct {
}

func (its *ILoad1) FetchOperand(reader *base.BytecodeReader) {
}

func (its *ILoad1) Execute(frame *rtdaBase.JavaFrame) {
	value := frame.LocalVars().GetInt(1)
	mylog.Printf("Value:%v", value)
	frame.OperandStack().PushInt(value)
}
