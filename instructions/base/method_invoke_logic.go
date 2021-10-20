package base

import (
	"fmt"

	rtdaBase "github.com/taoistwar/go-jvm/rtda/base"
	"github.com/taoistwar/go-jvm/rtda/java"
)

func InvokeMethod(invokerFrame *rtdaBase.JavaFrame, method *java.JavaMethod) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewJavaFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgSlotCount()) // 8字节类型占用2个Slot, 即1个long占用2个slot
	// 获取参数
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().ThisClassName(), method.Name(), method.Descriptor()))
		}
	}
}
