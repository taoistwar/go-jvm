package java

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (its *MethodDescriptor) addParameterType(t string) {
	pLen := len(its.parameterTypes)
	if pLen == cap(its.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, its.parameterTypes)
		its.parameterTypes = s
	}

	its.parameterTypes = append(its.parameterTypes, t)
}
