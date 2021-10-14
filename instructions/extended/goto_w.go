package extended

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type GotoW struct {
	offset int
}

func (its *GotoW) FetchOperand(reader *base.BytecodeReader) {
	its.offset = int(reader.ReadInt32())
}

func (its *GotoW) Execute(frame *rtdaBase.JavaFrame) {
	pc := frame.Thread().PC()
	nextPC := pc + its.offset
	frame.SetNextPC(nextPC)
}
