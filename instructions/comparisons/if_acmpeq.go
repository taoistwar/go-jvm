package comparisons

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type If_ACmpEQ struct {
	offset int
}

func (its *If_ACmpEQ) FetchOperand(reader *base.BytecodeReader) {
	its.offset = int(reader.ReadInt16())
}

func (its *If_ACmpEQ) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	val1 := stack.PopRef()
	val2 := stack.PopRef()
	if val1 == val2 { // TODO eq
		pc := frame.Thread().PC()
		nextPC := pc + its.offset
		frame.SetNextPC(nextPC)
	}
}
