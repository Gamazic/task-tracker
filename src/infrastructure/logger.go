package infrastructure

import (
	"fmt"
	"time"
)

type PrintLogger struct{}

func (PrintLogger) Errorf(template string, args ...any) {
	fmt.Print(time.Now().Format(time.DateTime), " | ")
	fmt.Printf(template, args...)
	fmt.Println()
}

func (PrintLogger) Infof(template string, args ...any) {
	fmt.Print(time.Now().Format(time.DateTime), " | ")
	fmt.Printf(template, args...)
	fmt.Println()
}

func (p PrintLogger) Error(msg string) {
	p.Errorf(msg)
}

func (p PrintLogger) LogIfErr(err error) {
	if err != nil {
		p.Error(err.Error())
	}
}
