package log

type Logger interface {
	PlainTextLogger
	StructuredLogger
	Tracer
}
