package extended

import (
	instBase "github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/instructions/loads"
	"github.com/taoistwar/go-jvm/instructions/math"
	"github.com/taoistwar/go-jvm/instructions/stores"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
)

type Wide struct {
	modifiedInstruction instBase.Instruction
}

func (its *Wide) FetchOperand(reader *instBase.BytecodeReader) {
	opcode := reader.ReadOperandCode()
	switch opcode {
	case 0x15:
		inst := &loads.ILoad{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLoad{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLoad{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLoad{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALoad{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x36:
		inst := &stores.IStore{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LStore{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FStore{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DStore{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.AStore{}
		inst.Index = uint(reader.ReadUint16())
		its.modifiedInstruction = inst
	case 0x84:
		inst := &math.IInc{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		its.modifiedInstruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (its *Wide) Execute(frame *rtdaBase.JavaFrame) {
	its.modifiedInstruction.Execute(frame)
}
