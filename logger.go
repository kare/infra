package infra

type Logger interface {
	Printf(format string, v ...any)
}
