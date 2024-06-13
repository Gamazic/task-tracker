package task

import (
	"fmt"
	"tracker_backend/internal/domain"
)

var ErrInvalidStage = fmt.Errorf("%w: invalid stage", domain.ErrDomain)

type Stage string

const (
	ToDo       = Stage("todo")
	InProgress = Stage("in_progress")
	Done       = Stage("done")
	Closed     = Stage("closed")
)

func (s Stage) Validate() error {
	switch s {
	case ToDo:
		return nil
	case InProgress:
		return nil
	case Done:
		return nil
	case Closed:
		return nil
	default:
		return ErrInvalidStage
	}
}
