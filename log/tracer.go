package log

import (
	"time"
)

type Tracer interface {
	TraceBegin(keysAndValues ...interface{}) TracerNeedATime
	TraceEnd(state interface{}, message ...string) TracerNeedATime
	TraceEndRoot(state interface{}, message ...string) TracerNeedATime
}

type TracerNeedATime interface {
	WithTime(moment time.Time)
}
