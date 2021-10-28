package java

import (
	"fmt"

	"github.com/taoistwar/go-jvm/classfile"
	"github.com/taoistwar/go-jvm/classpath"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type JavaClassLoader struct {
	classpath   *classpath.Classpath
	classMap    map[string]*JavaClass // loaded classes
	verboseFlag bool
}

func NewJClassLoader(cp *classpath.Classpath, verboseFlag bool) *JavaClassLoader {
	return &JavaClassLoader{
		classpath:   cp,
		classMap:    make(map[string]*JavaClass),
		verboseFlag: verboseFlag,
	}
}

func (its *JavaClassLoader) LoadJClass(name string) *JavaClass {
	if class, ok := its.classMap[name]; ok {
		// already loaded
		return class
	}

	var class *JavaClass
	if name[0] == '[' { // array class
		class = its.loadArrayClass(name)
	} else {
		class = its.loadNonArrayJavaClass(name)
	}

	if jlClassClass, ok := its.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}

	return class
}
func (its *JavaClassLoader) loadArrayClass(name string) *JavaClass {
	class := &JavaClass{
		accessFlags:   ACC_PUBLIC, // todo
		thisClassName: name,
		loader:        its,
		initStarted:   true,
	}
	class.superClass = its.LoadJClass("java/lang/Object")
	class.interfaces = []*JavaClass{its.LoadJClass("java/lang/Cloneable"), its.LoadJClass("java/io/Serializable")}
	its.classMap[name] = class
	return class
}
func (its *JavaClassLoader) loadNonArrayJavaClass(name string) *JavaClass {
	data, entry := its.readClass(name)
	class := its.defineJavaClass(data)
	link(class)
	if its.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (its *JavaClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := its.classpath.ReadClass(name)
	if err != nil {
		panic("java.lang.JavaClassNotFoundException: " + name)
	}
	if data == nil {
		panic("java.lang.JavaClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
func (its *JavaClassLoader) defineJavaClass(data []byte) *JavaClass {
	class := parseJavaClass(data)
	class.loader = its
	resolveSuperClass(class)
	resolveInterfaces(class)
	its.classMap[class.thisClassName] = class
	return class
}

func parseJavaClass(data []byte) *JavaClass {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.JavaClassFormatError")
		panic(err)
	}
	return NewJavaClass(cf)
}

// jvms 5.4.3.1
func resolveSuperClass(class *JavaClass) {
	if class.thisClassName != "java/lang/Object" {
		class.superClass = class.loader.LoadJClass(class.superClassName)
	}
}
func resolveInterfaces(class *JavaClass) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*JavaClass, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadJClass(interfaceName)
		}
	}
}

func link(class *JavaClass) {
	verify(class)
	prepare(class)
}

func verify(class *JavaClass) {
	// todo
}

// jvms 5.4.2
func prepare(class *JavaClass) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *JavaClass) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *JavaClass) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *JavaClass) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *JavaClass, field *JavaField) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JStringObject(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}
