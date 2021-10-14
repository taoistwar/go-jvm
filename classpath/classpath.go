package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func ParseClasspath(jreOption string, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return ".jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder")
}
func (classpath *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreDir = "/Users/taoistwar/Library/Java/JavaVirtualMachines/azul-1.8.0_292/Contents/Home" // for jdk 8
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "jre", "lib", "*")
	classpath.bootClasspath = newWillcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "jre", "lib", "ext", "*")
	classpath.extClasspath = newWillcardEntry(jreExtPath)
}
func (classpath *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	classpath.userClasspath = newEntry(cpOption)
}

func (its *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := its.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := its.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	return its.userClasspath.readClass(className)
}

func (classpath *Classpath) String() string {
	return classpath.userClasspath.String()
}
