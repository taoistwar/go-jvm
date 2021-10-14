package control

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type Goto struct {
	Offset int
}

func (its *Goto) FetchOperand(reader *base.BytecodeReader) {
	its.Offset = int(reader.ReadInt16())
}

func (its *Goto) Execute(frame *rtdaBase.JavaFrame) {
	pc := frame.Thread().PC()
	nextPC := pc + its.Offset
	frame.SetNextPC(nextPC)
}
