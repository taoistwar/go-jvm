package interpreter

import (
	"fmt"

	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/instructions/factory"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

func Interpret(javaMethod *java.JavaMethod, logInst bool) {

	thread := rtdaBase.NewJavaThread()
	frame := thread.NewJavaFrame(javaMethod)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, logInst)
}

func logFrames(thread *rtdaBase.JavaThread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().ThisClassName()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func catchErr(frame *rtdaBase.JavaFrame) {
	if r := recover(); r != nil {
		logFrames(frame.Thread())
		panic(r)
	}
}

func loop(thread *rtdaBase.JavaThread, logInst bool) {

	reader := &base.BytecodeReader{}

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadOperandCode()
		inst := factory.NewInstruction(opcode)
		inst.FetchOperand(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			// TODO thread close
			break
		}
	}
}
func logInstruction(frame *rtdaBase.JavaFrame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().ThisClassName()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
