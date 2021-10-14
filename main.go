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
	class := "com.github.taoistwar.java.MyObject"
	cpOption := "demo"
	// cp := classpath.ParseClasspath(cmd.XJreOption(), cmd.CpOption())
	cp := classpath.ParseClasspath(cmd.XJreOption(), cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cp.String(), class, cmd.Args())
	classLoader := java.NewJavaClassLoader(cp)
	className := strings.Replace(class, ".", "/", -1)

	mainClass := classLoader.LoadClass(className)
	if mainClass == nil {
		panic(fmt.Sprintf("Could not found or load main class %s\n", className))
	}
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		panic(fmt.Sprintf("Could not found main method in class %s\n", className))
	}
	interpreter.Interpret(mainMethod)
}

func main() {
	cmd := cli.ParseCmd()
	if cmd.VersionFlag() {
		println(version)
		return
	} else if cmd.HelpFlag() && cmd.Class() == "" {
		cli.PrintUsage()
	} else {
		startJvm(cmd)
	}
}
