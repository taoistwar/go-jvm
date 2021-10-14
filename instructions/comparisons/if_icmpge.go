package comparisons

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type If_ICmpGE struct {
	Offset int
}

func (its *If_ICmpGE) FetchOperand(reader *base.BytecodeReader) {
	its.Offset = int(reader.ReadInt16())
}

func (its *If_ICmpGE) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		pc := frame.Thread().PC()
		nextPC := pc + its.Offset
		frame.SetNextPC(nextPC)
	}
}
