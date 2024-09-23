package logging

import (
	"bytes"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/rs/zerolog"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var LogJSONLoggerEventHook zerolog.HookFunc = func(e *zerolog.Event, l zerolog.Level, msg string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(5, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	e.Str("func", frame.Function[strings.LastIndex(frame.Function, ".")+1:])
	e.Str("goroutine", string(bytes.Fields(debug.Stack())[1]))
}

var ContextLoggerEventHook zerolog.HookFunc = func(e *zerolog.Event, l zerolog.Level, msg string) {
	if spanContext := oteltrace.SpanContextFromContext(e.GetCtx()); spanContext.IsValid() {
		e.Str("spanID", spanContext.SpanID().String())
		e.Str("traceID", spanContext.TraceID().String())
	}
}
