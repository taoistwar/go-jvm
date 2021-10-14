package comparisons

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type If_ICmpGT struct {
	offset int
}

func (its *If_ICmpGT) FetchOperand(reader *base.BytecodeReader) {
	its.offset = int(reader.ReadInt16())
}

func (its *If_ICmpGT) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		pc := frame.Thread().PC()
		nextPC := pc + its.offset
		frame.SetNextPC(nextPC)
	}
}
