package java

import "strings"

type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (its *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	its.raw = descriptor
	its.parsed = &MethodDescriptor{}
	its.startParams()
	its.parseParamTypes()
	its.endParams()
	its.parseReturnType()
	its.finish()
	return its.parsed
}

func (its *MethodDescriptorParser) startParams() {
	if its.readUint8() != '(' {
		its.causePanic()
	}
}
func (its *MethodDescriptorParser) endParams() {
	if its.readUint8() != ')' {
		its.causePanic()
	}
}
func (its *MethodDescriptorParser) finish() {
	if its.offset != len(its.raw) {
		its.causePanic()
	}
}

func (its *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + its.raw)
}

func (its *MethodDescriptorParser) readUint8() uint8 {
	b := its.raw[its.offset]
	its.offset++
	return b
}
func (its *MethodDescriptorParser) unreadUint8() {
	its.offset--
}

func (its *MethodDescriptorParser) parseParamTypes() {
	for {
		t := its.parseFieldType()
		if t != "" {
			its.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (its *MethodDescriptorParser) parseReturnType() {
	if its.readUint8() == 'V' {
		its.parsed.returnType = "V"
		return
	}

	its.unreadUint8()
	t := its.parseFieldType()
	if t != "" {
		its.parsed.returnType = t
		return
	}

	its.causePanic()
}

func (its *MethodDescriptorParser) parseFieldType() string {
	switch its.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return its.parseObjectType()
	case '[':
		return its.parseArrayType()
	default:
		its.unreadUint8()
		return ""
	}
}

func (its *MethodDescriptorParser) parseObjectType() string {
	unread := its.raw[its.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		its.causePanic()
		return ""
	} else {
		objStart := its.offset - 1
		objEnd := its.offset + semicolonIndex + 1
		its.offset = objEnd
		descriptor := its.raw[objStart:objEnd]
		return descriptor
	}
}

func (its *MethodDescriptorParser) parseArrayType() string {
	arrStart := its.offset - 1
	its.parseFieldType()
	arrEnd := its.offset
	descriptor := its.raw[arrStart:arrEnd]
	return descriptor
}
