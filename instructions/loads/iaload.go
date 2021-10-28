package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IaLoad struct {
}

func (its *IaLoad) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IaLoad) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Ints()
	checkIndex(len(refs), index)
	stack.PushInt(refs[index])
}
