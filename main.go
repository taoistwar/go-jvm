package main

import (
	"fmt"
	"strings"

	"github.com/taoistwar/go-jvm/classfile"
	"github.com/taoistwar/go-jvm/classpath"
	"github.com/taoistwar/go-jvm/cli"
)

const version = "0.0.1"

func startJVM(cmd *cli.Cmd) {
	cp := classpath.Parse(cmd.XJreOption(), cmd.CpOption())
	fmt.Printf("classpath:%s class:%s args:%v\n", cp.String(), cmd.Class(), cmd.Args())

	className := strings.Replace(cmd.Class(), ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not found or load main class %s\n", cmd.Class())
		return
	}
	fmt.Printf("class data:%v\n", classData)
	cf, err := classfile.Parse(classData)
	if err != nil {
		fmt.Printf("Could not parse class file")
	}
	fmt.Printf("%v", cf)
}

func loadClass(className string, cp *classpath.Classpath) (cf *classfile.Classfile, err error) {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		msg := fmt.Sprintf("Could not found or load main class %s, %v\n", className, err)
		panic(msg)
	}
	fmt.Printf("class data:%v\n", classData)
	cf, err = classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return
}

func printClassInfo(cf *classfile.Classfile) {
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

func startJVMDev(cmd *cli.Cmd) {
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

func main() {
	cmd := cli.ParseCmd()
	if cmd.VersionFlag() {
		println(version)
		return
	} else if cmd.HelpFlag() && cmd.Class() == "" {
		cli.PrintUsage()
	} else {
		startJVMDev(cmd)
	}
}
