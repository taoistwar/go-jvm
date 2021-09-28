package cli

import (
	"flag"
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
	xJreOption  string
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help mesage")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.xJreOption, "XJre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func (cmd *Cmd) HelpFlag() bool {
	return cmd.helpFlag
}

func (cmd *Cmd) VersionFlag() bool {
	return cmd.versionFlag
}

func (cmd *Cmd) CpOption() string {
	return cmd.cpOption
}

func (cmd *Cmd) Class() string {
	return cmd.class
}

func (cmd *Cmd) Args() []string {
	return cmd.args
}
func (cmd *Cmd) XJreOption() string {
	return cmd.xJreOption
}
