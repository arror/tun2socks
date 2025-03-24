package log

/*
#include <stdlib.h>

typedef enum {
    LOG_APP = 0,
} LogType;

typedef enum {
    LOG_DEBUG = 3,
    LOG_INFO = 4,
    LOG_WARN = 5,
    LOG_ERROR = 6,
    LOG_FATAL = 7,
} LogLevel;

static void OHOS_LOG(LogType type, LogLevel level, unsigned int domain, const char *tag, const char *message) {
    extern int OH_LOG_Print(LogType type, LogLevel level, unsigned int domain, const char *tag, const char *fmt, ...);
    OH_LOG_Print(type, level, domain, tag, "%{public}s", message);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var ohosLog = func(level C.LogLevel, message string) {
	temp := C.CString(message)
	tag := C.CString("GOLANG")
	defer C.free(unsafe.Pointer(temp))
	defer C.free(unsafe.Pointer(tag))
	C.OHOS_LOG(C.LOG_APP, level, 0x0F, tag, temp)
}

func Debug(message string) {
	ohosLog(C.LOG_DEBUG, message)
}

func Info(message string) {
	ohosLog(C.LOG_INFO, message)
}

func Warn(message string) {
	ohosLog(C.LOG_WARN, message)
}

func Error(message string) {
	ohosLog(C.LOG_ERROR, message)
}

func Fatal(message string) {
	ohosLog(C.LOG_FATAL, message)
}

func Debugf(format string, a ...any) {
	Debug(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) {
	Info(fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...any) {
	Warn(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...any) {
	Error(fmt.Sprintf(format, a...))
}

func Fatalf(format string, a ...any) {
	Fatal(fmt.Sprintf(format, a...))
}
