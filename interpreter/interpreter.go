package interpreter

import (
	"encoding/json"
	"fmt"

	"github.com/taoistwar/go-jvm/instructions/base"
	"github.com/taoistwar/go-jvm/instructions/factory"
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

func Interpret(javaMethod *java.JavaMethod) {

	thread := rtdaBase.NewJavaThread()
	frame := thread.NewJavaFrame(javaMethod)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, javaMethod.Code())
}

func catchErr(frame *rtdaBase.JavaFrame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtdaBase.JavaThread, bytecode []byte) {

	reader := &base.BytecodeReader{}

	reader.ResetCode(bytecode)
	for {
		frame := thread.TopFrame()
		if frame == nil {
			return
		}
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.ResetPC(pc)
		opcode := reader.ReadOperandCode()
		inst := factory.NewInstruction(opcode)
		inst.FetchOperand(reader)
		frame.SetNextPC(reader.PC())

		// execute
		data, err := json.Marshal(inst)
		if err != nil {
			fmt.Printf("\npc: %2d inst: %T %v\n", pc, inst, inst)
		} else {
			fmt.Printf("\npc: %2d inst: %T %v\n", pc, inst, string(data))
		}
		inst.Execute(frame)
	}
}
