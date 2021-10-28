package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type LaLoad struct {
}

func (its *LaLoad) FetchOperand(reader *base.BytecodeReader) {

}

func (its *LaLoad) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Longs()
	checkIndex(len(refs), index)
	stack.PushLong(refs[index])
}
