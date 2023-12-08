package task

import (
	"fmt"
	"tracker_backend/src/domain"
)

var ErrInvalidTaskNumber = fmt.Errorf("%w: invalid task number", domain.ErrDomain)

type TaskNumber int

func (t TaskNumber) Validate() error {
	if t < 0 {
		return ErrInvalidTaskNumber
	}
	return nil
}
