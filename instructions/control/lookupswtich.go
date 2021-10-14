package control

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LookupSwitch struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (self *LookupSwitch) FetchOperand(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LookupSwitch) Execute(frame *rtdaBase.JavaFrame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			pc := frame.Thread().PC()
			nextPC := pc + int(offset)
			frame.SetNextPC(nextPC)
			return
		}
	}

	pc := frame.Thread().PC()
	nextPC := pc + int(self.defaultOffset)
	frame.SetNextPC(nextPC)
}
