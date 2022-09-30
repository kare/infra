package infra

// Logger implements a common logger interaface compatible with standard [log.Logger].
type Logger interface {
	Printf(format string, v ...any)
}
