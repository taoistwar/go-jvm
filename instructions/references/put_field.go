package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type PutField struct {
	Index uint
}

func (its *PutField) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}

func (its *PutField) Execute(frame *rtdaBase.JavaFrame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(its.Index).(*java.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		mylog.Printf("Index:%v, Field:%v %v, Slot:%v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), val, ref)
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		mylog.Printf("Index:%v, Field:%v %v, Slot:%v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), val, ref)
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		mylog.Printf("Index:%v, Field:%v %v, Slot:%v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), val, ref)
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		mylog.Printf("Index:%v, Field:%v %v, Slot:%v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), val, ref)
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		mylog.Printf("Index:%v, Field:%v %v, Slot:%v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), val, ref)
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}

}
