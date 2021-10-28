package java

type JavaObject struct {
	class *JavaClass
	data  interface{}
	extra interface{}
}

func newJavaObject(class *JavaClass) *JavaObject {
	return &JavaObject{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (its *JavaObject) Extra() interface{} {
	return its.extra
}
func (its *JavaObject) SetExtra(extra interface{}) {
	its.extra = extra
}
func (its *JavaObject) Class() *JavaClass {
	return its.class
}
func (its *JavaObject) Fields() Slots {
	return its.data.(Slots)
}
func (its *JavaObject) Bytes() []int8 {
	return its.data.([]int8)
}
func (its *JavaObject) Shorts() []int16 {
	return its.data.([]int16)
}
func (its *JavaObject) Ints() []int32 {
	return its.data.([]int32)
}
func (its *JavaObject) Longs() []int64 {
	return its.data.([]int64)
}
func (its *JavaObject) Chars() []uint16 {
	return its.data.([]uint16)
}
func (its *JavaObject) Floats() []float32 {
	return its.data.([]float32)
}
func (its *JavaObject) Doubles() []float64 {
	return its.data.([]float64)
}
func (its *JavaObject) Refs() []*JavaObject {
	return its.data.([]*JavaObject)
}
func (its *JavaObject) ArrayLength() int32 {
	switch data := its.data.(type) {
	case []int8:
		return int32(len(data))
	case []int16:
		return int32(len(data))
	case []int32:
		return int32(len(data))
	case []int64:
		return int32(len(data))
	case []uint16:
		return int32(len(data))
	case []float32:
		return int32(len(data))
	case []float64:
		return int32(len(data))
	case []*JavaObject:
		return int32(len(data))
	default:
		panic("Not array")
	}
}

func (its *JavaObject) IsInstanceOf(class *JavaClass) bool {
	return class.IsAssignableFrom(its.class)
}

// reflection
func (self *JavaObject) GetRefVar(name, descriptor string) *JavaObject {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (self *JavaObject) SetRefVar(name, descriptor string, ref *JavaObject) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
func (self *JavaObject) SetIntVar(name, descriptor string, val int32) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetInt(field.slotId, val)
}
func (self *JavaObject) GetIntVar(name, descriptor string) int32 {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetInt(field.slotId)
}
