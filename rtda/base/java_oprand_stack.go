package base

import (
	"math"

	"github.com/taoistwar/go-jvm/rtda/java"
)

type JavaOperandStack struct {
	size  uint
	slots []JavaLocalVarSlot
}

func newJavaOperandStack(maxStack uint) *JavaOperandStack {
	if maxStack > 0 {
		return &JavaOperandStack{
			size:  0,
			slots: make([]JavaLocalVarSlot, maxStack),
		}
	}
	return nil
}

func (its *JavaOperandStack) PushInt(value int32) {
	its.slots[its.size].Data = value
	its.size++
}

func (its *JavaOperandStack) PopInt() int32 {
	its.size--
	return its.slots[its.size].Data
}

func (jos *JavaOperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	jos.slots[jos.size].Data = int32(bits)
	jos.size++
}

func (jos *JavaOperandStack) PopFloat() float32 {
	jos.size--
	bits := uint32(jos.slots[jos.size].Data)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (jos *JavaOperandStack) PushLong(val int64) {
	jos.slots[jos.size].Data = int32(val)
	jos.slots[jos.size+1].Data = int32(val >> 32)
	jos.size += 2
}

func (jos *JavaOperandStack) PopLong() int64 {
	jos.size -= 2
	low := uint32(jos.slots[jos.size].Data)
	high := uint32(jos.slots[jos.size+1].Data)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (jos *JavaOperandStack) PushDouble(value float64) {
	bits := math.Float64bits(value)
	jos.PushLong(int64(bits))
}

func (jos *JavaOperandStack) PopDouble() float64 {
	bits := uint64(jos.PopLong())
	return math.Float64frombits(bits)
}

func (jos *JavaOperandStack) PushRef(ref *java.JavaObject) {
	jos.slots[jos.size].ref = ref
	jos.size++
}

func (jos *JavaOperandStack) PopRef() *java.JavaObject {
	jos.size--
	ref := jos.slots[jos.size].ref
	jos.slots[jos.size].ref = nil
	return ref
}
func (jos *JavaOperandStack) PushSlot(slot *JavaLocalVarSlot) {
	jos.slots[jos.size] = *slot
	jos.size++
}

func (jos *JavaOperandStack) PopSlot() *JavaLocalVarSlot {
	jos.size--
	return &jos.slots[jos.size]
}

func (its *JavaOperandStack) GetRefFromTop(n uint) *java.JavaObject {
	return its.slots[its.size-1-n].ref
}
