package references

import (
	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/mylog"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

type PutStatic struct {
	Index uint
}

func (its *PutStatic) FetchOperand(reader *base.BytecodeReader) {
	its.Index = uint(reader.ReadInt16())
}

func (its *PutStatic) Execute(frame *rtdaBase.JavaFrame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	ref := cp.GetConstant(its.Index)
	switch fieldRef := ref.(type) {
	case *java.FieldRef:
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
		if field.IsFinal() {
			if currentClass != class || currentMethod.Name() != "<clinit>" {
				panic("java.lang.IllegalAccessError")
			}
		}

		descriptor := field.Descriptor()
		slotId := field.SlotId()
		slots := class.StaticVars()
		stack := frame.OperandStack()

		switch descriptor[0] {
		case 'Z', 'B', 'C', 'S', 'I':
			value := stack.PopInt()
			mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
			slots.SetInt(slotId, value)
		case 'F':
			value := stack.PopFloat()
			mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
			slots.SetFloat(slotId, value)
		case 'J':
			value := stack.PopLong()
			mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
			slots.SetLong(slotId, value)
		case 'D':
			value := stack.PopDouble()
			mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
			slots.SetDouble(slotId, value)
		case 'L', '[':
			value := stack.PopRef()
			mylog.Printf("Index: %v, Field: %v %v, Value:%v", its.Index, fieldRef.Name(), fieldRef.Descriptor(), value)
			slots.SetRef(slotId, value)
		default:
			// todo
		}
	default:
		panic("put_static index not field ref")

	}

}
