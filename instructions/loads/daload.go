package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type DaLoad struct {
}

func (its *DaLoad) FetchOperand(reader *base.BytecodeReader) {
}

func (its *DaLoad) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Doubles()
	checkIndex(len(refs), index)
	stack.PushDouble(refs[index])
}
