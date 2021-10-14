package java

type JavaObject struct {
	class  *JavaClass
	fields Slots
}

func newJavaObject(class *JavaClass) *JavaObject {
	return &JavaObject{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (jo *JavaObject) Class() *JavaClass {
	return jo.class
}
func (jo *JavaObject) Fields() Slots {
	return jo.fields
}

func (jo *JavaObject) IsInstanceOf(class *JavaClass) bool {
	return class.isAssignableFrom(jo.class)
}
