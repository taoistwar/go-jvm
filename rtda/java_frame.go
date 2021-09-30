package rtda

type JavaFrame struct {
	lower        *JavaFrame
	localVars    JavaLocalVars
	operandStack *JavaOperandStack
}

func NewJavaFrame(maxLocalVars uint, maxOperandStack uint) *JavaFrame {
	return &JavaFrame{
		localVars:    newJavaLocalVars(maxLocalVars),
		operandStack: newJavaOperandStack(maxOperandStack),
	}
}
