package base

import (
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type Instruction interface {
	// 从字节码中提取操作数
	FetchOperand(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *rtdaBase.JavaFrame)
}

type BranchInstruction struct {
	Offset int
}

func (its *BranchInstruction) FetchOperand(reader *BytecodeReader) {
	its.Offset = int(reader.ReadInt16())
}
func (its *BranchInstruction) Execute(frame *rtdaBase.JavaFrame) {

}

type Index8Instruction struct {
	Index uint
}

func (its *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	its.Index = uint(reader.ReadOperandCode())
}

type Index16Instruction struct {
	Index uint
}

func (its *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	its.Index = uint(reader.ReadUint16())
}
