package microframework

type Logger interface {
	Errorf(template string, args ...any)
	Infof(template string, args ...any)
	Error(msg string)
	LogIfErr(err error)
}
