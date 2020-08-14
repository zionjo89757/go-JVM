package reserved

import "jvm/instructions/base"
import "jvm/rtda"
import "jvm/native"
import _ "jvm/native/java/io"
import _ "jvm/native/java/lang"
import _ "jvm/native/java/security"
import _ "jvm/native/java/util/concurrent/atomic"
import _ "jvm/native/sun/io"
import _ "jvm/native/sun/misc"
import _ "jvm/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
