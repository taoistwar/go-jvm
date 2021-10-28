package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type CaLoad struct {
}

func (its *CaLoad) FetchOperand(reader *base.BytecodeReader) {
}

func (its *CaLoad) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Chars()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}
