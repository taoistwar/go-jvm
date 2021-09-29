package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []*AttributeInfo {
	attributesCount := reader.ReadUint16()
	attributes := make([]*AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) *AttributeInfo {
	attrNameIndex := reader.ReadUint16()
	attrLen := reader.ReadUint32()
	attrName := cp.getUtf8(attrNameIndex)
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return &attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	// TODO jdk 17 has more type
	switch attrName {
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []*AttributeInfo
}

func (its *CodeAttribute) readInfo(reader *ClassReader) {
	its.maxStack = reader.ReadUint16()
	its.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	its.code = reader.ReadBytes(codeLength)
	its.exceptionTable = readExceptionTable(reader)
	its.attributes = readAttributes(reader, its.cp)
}

func (its *CodeAttribute) MaxStack() uint {
	return uint(its.maxStack)
}
func (its *CodeAttribute) MaxLocals() uint {
	return uint(its.maxLocals)
}
func (its *CodeAttribute) Code() []byte {
	return its.code
}
func (its *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return its.exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.ReadUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}
	return exceptionTable
}

func (its *ExceptionTableEntry) StartPc() uint16 {
	return its.startPc
}
func (its *ExceptionTableEntry) EndPc() uint16 {
	return its.endPc
}
func (its *ExceptionTableEntry) HandlerPc() uint16 {
	return its.handlerPc
}
func (its *ExceptionTableEntry) CatchType() uint16 {
	return its.catchType
}

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (its *ConstantValueAttribute) readInfo(reader *ClassReader) {
	its.constantValueIndex = reader.ReadUint16()
}

func (its *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return its.constantValueIndex
}

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (its *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (its *ExceptionsAttribute) readInfo(reader *ClassReader) {
	its.exceptionIndexTable = reader.ReadUint16s()
}

func (its *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return its.exceptionIndexTable
}

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (its *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	its.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range its.lineNumberTable {
		its.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.ReadUint16(),
			lineNumber: reader.ReadUint16(),
		}
	}
}

func (its *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(its.lineNumberTable) - 1; i >= 0; i-- {
		entry := its.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (its *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.ReadUint16()
	its.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range its.localVariableTable {
		its.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.ReadUint16(),
			length:          reader.ReadUint16(),
			nameIndex:       reader.ReadUint16(),
			descriptorIndex: reader.ReadUint16(),
			index:           reader.ReadUint16(),
		}
	}
}

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (its *SourceFileAttribute) readInfo(reader *ClassReader) {
	its.sourceFileIndex = reader.ReadUint16()
}

func (its *SourceFileAttribute) FileName() string {
	return its.cp.getUtf8(its.sourceFileIndex)
}

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (its *UnparsedAttribute) readInfo(reader *ClassReader) {
	its.info = reader.ReadBytes(its.length)
}

func (its *UnparsedAttribute) Info() []byte {
	return its.info
}
