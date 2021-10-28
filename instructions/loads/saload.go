package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type SaLoad struct {
}

func (its *SaLoad) FetchOperand(reader *base.BytecodeReader) {

}

func (its *SaLoad) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Shorts()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}
