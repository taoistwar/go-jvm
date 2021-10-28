package loads

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type FaLoad struct {
}

func (its *FaLoad) FetchOperand(reader *base.BytecodeReader) {

}

func (its *FaLoad) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Floats()
	checkIndex(len(refs), index)
	stack.PushFloat(refs[index])
}
