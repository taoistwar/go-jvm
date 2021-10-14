package java

import "math"

type Slot struct {
	value int32
	ref   *JavaObject
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (slots Slots) SetInt(index uint, value int32) {
	slots[index].value = value
}
func (slots Slots) GetInt(index uint) int32 {
	return slots[index].value
}

func (slots Slots) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	slots[index].value = int32(bits)
}
func (slots Slots) GetFloat(index uint) float32 {
	bits := uint32(slots[index].value)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (slots Slots) SetLong(index uint, value int64) {
	slots[index].value = int32(value)
	slots[index+1].value = int32(value >> 32)
}
func (slots Slots) GetLong(index uint) int64 {
	low := uint32(slots[index].value)
	high := uint32(slots[index+1].value)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (slots Slots) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	slots.SetLong(index, int64(bits))
}
func (slots Slots) GetDouble(index uint) float64 {
	bits := uint64(slots.GetLong(index))
	return math.Float64frombits(bits)
}

func (slots Slots) SetRef(index uint, ref *JavaObject) {
	slots[index].ref = ref
}
func (slots Slots) GetRef(index uint) *JavaObject {
	return slots[index].ref
}
