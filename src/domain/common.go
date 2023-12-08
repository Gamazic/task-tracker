package domain

import "errors"

var ErrDomain = errors.New("domain error")

type Validatable interface {
	Validate() error
}

func ValidateCascade(validatable []Validatable) error {
	var err error
	for i := range validatable {
		err = validatable[i].Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
