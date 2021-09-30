package classfile

import (
	"fmt"
	"strings"

	"github.com/taoistwar/go-jvm/classpath"
	"github.com/taoistwar/go-jvm/cli"
)

func Parse(data []byte) (cf *Classfile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{data: data}
	cf = &Classfile{}
	cf.read(cr)
	return
}

func startJvmByClassfile(cmd *cli.Cmd) {
	class := "com.github.taoistwar.java.ClassFileTest"
	cpOption := "demo"
	cp := classpath.Parse(cmd.XJreOption(), cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cp.String(), class, cmd.Args())
	className := strings.Replace(class, ".", "/", -1)

	cf, err := loadClass(className, cp)
	if err != nil {
		panic(err)
	}
	if cf.MajorVersion() > 55 && cf.MinorVersion() == 65535 && !cmd.XPreviewFeatureOption() {
		// 开启 preview featrue
		panic("classfile with preview feature, must enable preview with -XPreviewFeature")
	}
	printClassInfo(cf)

}

func loadClass(className string, cp *classpath.Classpath) (cf *Classfile, err error) {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		msg := fmt.Sprintf("Could not found or load main class %s, %v\n", className, err)
		panic(msg)
	}
	fmt.Printf("class data:%v\n", classData)
	cf, err = Parse(classData)
	if err != nil {
		panic(err)
	}
	return
}

func printClassInfo(cf *Classfile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
