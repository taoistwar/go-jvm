package control

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type TableSwitch struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (its *TableSwitch) FetchOperand(reader *base.BytecodeReader) {
	reader.SkipPadding()
	its.defaultOffset = reader.ReadInt32()
	its.low = reader.ReadInt32()
	its.high = reader.ReadInt32()
	jumpOffsetsCount := its.high - its.low + 1
	its.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (its *TableSwitch) Execute(frame *rtdaBase.JavaFrame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= its.low && index <= its.high {
		offset = int(its.jumpOffsets[index-its.low])
	} else {
		offset = int(its.defaultOffset)
	}

	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
