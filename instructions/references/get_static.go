package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type GetStatic struct {
	Index uint
}

func (its *GetStatic) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}

func (its *GetStatic) Execute(frame *rtdaBase.JavaFrame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(its.Index).(*java.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		value := slots.GetInt(slotId)
		mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
		stack.PushInt(value)
	case 'F':
		value := slots.GetFloat(slotId)
		mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
		stack.PushFloat(value)
	case 'J':
		value := slots.GetLong(slotId)
		mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
		stack.PushLong(value)
	case 'D':
		value := slots.GetDouble(slotId)
		mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
		stack.PushDouble(value)
	case 'L', '[':
		value := slots.GetRef(slotId)
		mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
		stack.PushRef(value)
	default:
		// todo
	}

}
