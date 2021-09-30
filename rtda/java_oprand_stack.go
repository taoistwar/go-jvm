package rtda

import "math"

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
	its.slots[its.size].data = value
	its.size++
}

func (its *JavaOperandStack) PopInt() int32 {
	its.size--
	return its.slots[its.size].data
}

func (self *JavaOperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].data = int32(bits)
	self.size++
}
func (self *JavaOperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].data)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (self *JavaOperandStack) PushLong(val int64) {
	self.slots[self.size].data = int32(val)
	self.slots[self.size+1].data = int32(val >> 32)
	self.size += 2
}
func (self *JavaOperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].data)
	high := uint32(self.slots[self.size+1].data)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (self *JavaOperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *JavaOperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *JavaOperandStack) PushRef(ref *JavaObject) {
	self.slots[self.size].ref = ref
	self.size++
}

func (self *JavaOperandStack) PopRef() *JavaObject {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}
