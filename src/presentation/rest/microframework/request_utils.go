package microframework

import (
	"encoding/json"
	"errors"
	"io"
)

var ValidationErr = errors.New("validation error")

type Validatable interface {
	Validate() error
}

func ReadValidate(r io.ReadCloser, v Validatable) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return err
	}
	err = v.Validate()
	if err != nil {
		return err
	}
	return nil
}
