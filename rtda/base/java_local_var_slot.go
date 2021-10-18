package base

import "github.com/taoistwar/go-jvm/rtda/java"

type JavaLocalVarSlot struct {
	data int32
	ref  *java.JavaObject
}

func (its *JavaLocalVarSlot) GetData() int32 {
	return its.data
}
func (its *JavaLocalVarSlot) SetData(data int32) {
	its.data = data
}

func (its *JavaLocalVarSlot) GetRef() *java.JavaObject {
	return its.ref
}
func (its *JavaLocalVarSlot) SetRef(ref *java.JavaObject) {
	its.ref = ref
}
