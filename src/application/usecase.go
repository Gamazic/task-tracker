package application

type EmptyOutputType struct{}

var EmptyOutput EmptyOutputType

type Usecase[InputDto any, OutputDto any] interface {
	Execute(InputDto) (OutputDto, error)
}
