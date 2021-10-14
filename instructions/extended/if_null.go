package extended

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IfNull struct {
	Offset int
}

func (its *IfNull) FetchOperand(reader *base.BytecodeReader) {
	its.Offset = int(reader.ReadInt16())
}

func (its *IfNull) Execute(frame *rtdaBase.JavaFrame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		pc := frame.Thread().PC()
		nextPC := pc + its.Offset
		frame.SetNextPC(nextPC)
	}
}
