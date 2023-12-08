package task

import (
	"fmt"
	"tracker_backend/src/domain"
)

var (
	ErrInvalidDescription = fmt.Errorf("%w: invalid description", domain.ErrDomain)
	ErrEmptyDescription   = fmt.Errorf("%w: empty description", ErrInvalidDescription)
	ErrDescriptionTooLong = fmt.Errorf("%w: description is too long", ErrInvalidDescription)
)

const maxDescriptionLen = 255

type Description string

func (d Description) Validate() error {
	if d == "" {
		return ErrEmptyDescription
	}
	if len(d) > maxDescriptionLen {
		return ErrDescriptionTooLong
	}
	return nil
}
