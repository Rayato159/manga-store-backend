package utils

import (
	"runtime"
	"time"
)

func Trace() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	if frames == nil {
		return "error, can't trace a function name"
	}
	frame, _ := frames.Next()
	return frame.Function
}

func CallTimer(start int64) int64 {
	return time.Now().UnixMilli() - start
}
