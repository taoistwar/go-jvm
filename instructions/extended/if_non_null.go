package extended

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IfNonNull struct {
	Offset int
}

func (its *IfNonNull) FetchOperand(reader *base.BytecodeReader) {
	its.Offset = int(reader.ReadInt16())
}

func (its *IfNonNull) Execute(frame *rtdaBase.JavaFrame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		pc := frame.Thread().PC()
		nextPC := pc + its.Offset
		frame.SetNextPC(nextPC)
	}
}
