package comparisons

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IfEQ struct {
	offset int
}

func (its *IfEQ) FetchOperand(reader *base.BytecodeReader) {
	its.offset = int(reader.ReadInt16())
}

func (its *IfEQ) Execute(frame *rtdaBase.JavaFrame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		pc := frame.Thread().PC()
		nextPC := pc + its.offset
		frame.SetNextPC(nextPC)
	}
}
