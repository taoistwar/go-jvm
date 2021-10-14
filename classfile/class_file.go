package classfile

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	magic         uint32
	minor_version uint16
	major_version uint16
	constant_pool ConstantPool
	access_flag   uint16
	this_class    uint16
	super_class   uint16
	interfaces    []uint16
	fields        []*MemberInfo
	methods       []*MemberInfo
	attributes    []*AttributeInfo
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constant_pool = readConstantPool(reader)
	cf.access_flag = reader.ReadUint16()
	cf.this_class = reader.ReadUint16()
	cf.super_class = reader.ReadUint16()
	cf.interfaces = reader.ReadUint16s()
	cf.fields = readMembers(reader, cf.constant_pool)
	cf.methods = readMembers(reader, cf.constant_pool)
	cf.attributes = readAttributes(reader, cf.constant_pool)
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.ReadUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// For a class file whose major_version is 56 or above, the minor_version must be 0 or 65535.
	// For a class file whose major_version is between 45 and 55 inclusive, the minor_version may be any value.
	cf.minor_version = reader.ReadUint16()
	cf.major_version = reader.ReadUint16()
	if cf.major_version >= 45 || cf.major_version <= 55 {
		return
	}
	if cf.major_version >= 55 {
		if cf.minor_version == 0 {
			return
		}
		if cf.minor_version == 65535 { // has preview features
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
func (cf *ClassFile) Magic() uint32 {
	return cf.magic
}
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minor_version
}
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.major_version
}
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constant_pool
}
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.access_flag
}
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constant_pool.getClassName(cf.this_class)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.super_class > 0 {
		return cf.constant_pool.getClassName(cf.super_class)
	}
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constant_pool.getClassName(cpIndex)
	}
	return interfaceNames
}
