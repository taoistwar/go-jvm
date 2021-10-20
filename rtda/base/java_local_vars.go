package base

import (
	"math"

	java "github.com/taoistwar/go-jvm/rtda/java"
)

type JavaLocalVars []JavaLocalVarSlot

func newJavaLocalVars(maxLocalVars uint) JavaLocalVars {
	if maxLocalVars > 0 {
		data := make([]JavaLocalVarSlot, maxLocalVars)
		return data
	}
	return nil
}

func (its JavaLocalVars) SetInt(index uint, value int32) {
	(its)[index].data = value
}
func (its JavaLocalVars) GetInt(index uint) int32 {
	return (its)[index].data
}

// long consumes two slots
func (its JavaLocalVars) SetLong(index uint, val int64) {
	its[index].data = int32(val)
	its[index+1].data = int32(val >> 32)
}
func (its JavaLocalVars) GetLong(index uint) int64 {
	low := uint32(its[index].data)
	high := uint32(its[index+1].data)
	return int64(high)<<32 | int64(low)
}

func (its JavaLocalVars) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	(its)[index+1].SetData(int32(bits))
}
func (its JavaLocalVars) GetFloat(index uint) float32 {
	bits := uint32((its)[index+1].data)
	return math.Float32frombits(bits)
}

// double consumes two slots
func (its JavaLocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	its.SetLong(index, int64(bits))
}
func (its JavaLocalVars) GetDouble(index uint) float64 {
	bits := uint64(its.GetLong(index))
	return math.Float64frombits(bits)
}

func (its JavaLocalVars) SetRef(index uint, ref *java.JavaObject) {
	(its)[index].ref = ref
}
func (its JavaLocalVars) GetRef(index uint) *java.JavaObject {
	return (its)[index].GetRef()
}

func (self JavaLocalVars) SetSlot(index uint, slot *JavaLocalVarSlot) {
	self[index] = *slot
}
