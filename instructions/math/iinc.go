package math

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type IInc struct {
	Index uint
	Const int32
}

func (its *IInc) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt8())
	its.Const = int32(reader.ReadInt8())
}

func (its *IInc) Execute(frame *rtdaBase.JavaFrame) {
	localVars := frame.LocalVars()
	value := localVars.GetInt(its.Index)
	value += its.Const
	localVars.SetInt(its.Index, value)
}
