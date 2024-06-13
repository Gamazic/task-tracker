package domain

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
