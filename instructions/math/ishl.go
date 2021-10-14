package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IShL struct {
}

func (its *IShL) FetchOperand(reader *base.BytecodeReader) {

}

func (its *IShL) Execute(frame *rtdaBase.JavaFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}
