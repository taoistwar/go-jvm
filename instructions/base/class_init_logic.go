package base

import (
	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

// jvms 5.5
func InitClass(thread *rtdaBase.JavaThread, class *java.JavaClass) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtdaBase.JavaThread, class *java.JavaClass) {
	init := class.GetClassInitMethod()
	if init != nil {
		// exec <clinit>
		newFrame := thread.NewJavaFrame(init)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtdaBase.JavaThread, class *java.JavaClass) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
