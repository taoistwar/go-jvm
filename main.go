package main

import (
	"fmt"
	"strings"

	"github.com/taoistwar/go-jvm/classfile"
	"github.com/taoistwar/go-jvm/classpath"
	"github.com/taoistwar/go-jvm/cli"
	"github.com/taoistwar/go-jvm/rtda"
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

func main() {
	cmd := cli.ParseCmd()
	if cmd.VersionFlag() {
		println(version)
		return
	} else if cmd.HelpFlag() && cmd.Class() == "" {
		cli.PrintUsage()
	} else {
		rtda.StartJvmByRtda(cmd)
	}
}
