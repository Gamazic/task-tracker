package task

import (
	"fmt"
	"tracker_backend/src/domain"
)

var ErrInvalidStage = fmt.Errorf("%w: invalid stage", domain.ErrDomain)

type Stage string

const (
	ToDo       = Stage("todo")
	InProgress = Stage("in_progress")
	Done       = Stage("done")
)

func (s Stage) Validate() error {
	switch s {
	case ToDo:
		return nil
	case InProgress:
		return nil
	case Done:
		return nil
	default:
		return ErrInvalidStage
	}
}
