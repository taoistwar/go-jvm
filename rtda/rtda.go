package rtda

import "github.com/taoistwar/go-jvm/cli"

func testLocalVars(vars JavaLocalVars) {
	vars.SetInt(0, 100)
	println(vars.GetInt(0))
	vars.SetInt(1, -100)
	println(vars.GetInt(1))
	vars.SetLong(2, 2997924580)
	println(vars.GetLong(2))
	vars.SetLong(4, -2997924580)
	println(vars.GetLong(4))
	vars.SetFloat(6, 3.1415926)
	println(vars.GetFloat(6))
	vars.SetDouble(7, 2.71828182845)
	println(vars.GetDouble(7))
	vars.SetRef(9, nil)
	println(vars.GetRef(9))

}
func testOperandStack(ops *JavaOperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
func StartJvmByRtda(cmd *cli.Cmd) {
	frame := NewJavaFrame(100, 100)
	testLocalVars(frame.localVars)
	testOperandStack(frame.operandStack)
}
