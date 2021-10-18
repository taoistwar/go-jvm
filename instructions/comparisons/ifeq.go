package comparisons

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IfEQ struct {
	Offset int
}

func (its *IfEQ) FetchOperand(reader *base.BytecodeReader) {
	its.Offset = int(reader.ReadInt16())
}

func (its *IfEQ) Execute(frame *rtdaBase.JavaFrame) {
	val := frame.OperandStack().PopInt()
	mylog.Printf("Offset: %v, Value:%v, Skip: %v", its.Offset, val, val == 0)
	if val == 0 {
		pc := frame.Thread().PC()
		nextPC := pc + its.Offset
		frame.SetNextPC(nextPC)
	}
}
