package main

import (
	"fmt"
	"strings"

	"github.com/taoistwar/go-jvm/classpath"
	"github.com/taoistwar/go-jvm/cli"
	"github.com/taoistwar/go-jvm/interpreter"
	"github.com/taoistwar/go-jvm/rtda/java"
)

const version = "0.0.1"

func startJvm(cmd *cli.Cmd) {
	cp := classpath.ParseClasspath(cmd.XJreOption(), cmd.CpOption())
	fmt.Printf("classpath:%s class:%s args:%v\n", cp.String(), cmd.Class(), cmd.Args())
	classLoader := java.NewJavaClassLoader(cp, cmd.VerboseClassFlag())
	className := strings.Replace(cmd.Class(), ".", "/", -1)

	mainClass := classLoader.LoadClass(className)
	if mainClass == nil {
		panic(fmt.Sprintf("Could not found or load main class %s\n", className))
	}
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		panic(fmt.Sprintf("Could not found main method in class %s\n", className))
	}
	interpreter.Interpret(mainMethod, cmd.VerboseInstFlag())
}

func main() {
	cmd := cli.ParseCmd()
	class := "com.github.taoistwar.java.InvokeDemo"
	cpOption := "demo"
	cmd.Reset(class, cpOption)
	if cmd.VersionFlag() {
		println(version)
		return
	} else if cmd.HelpFlag() && cmd.Class() == "" {
		cli.PrintUsage()
	} else {
		startJvm(cmd)
	}
}
