package infrastructure

import "fmt"

type Logger interface {
	Errorf(template string, args ...any)
	Error(msg string)
	LogIfErr(err error)
}

type PrintLogger struct{}

func (PrintLogger) Errorf(template string, args ...any) {
	fmt.Printf(template, args...)
	fmt.Println()
}

func (PrintLogger) Error(msg string) {
	fmt.Println(msg)
}

func (p PrintLogger) LogIfErr(err error) {
	if err != nil {
		p.Error(err.Error())
	}
}
