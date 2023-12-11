package register_controller

import (
	"fmt"
	"tracker_backend/src/presentation/rest/microframework"
)

type RegisterRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u RegisterRequestModel) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("%w: non empty field 'username' is required",
			microframework.ValidationErr)
	}
	if u.Password == "" {
		return fmt.Errorf("%w: non empty field 'password' is required",
			microframework.ValidationErr)
	}
	return nil
}

type RegisterResponseModel struct {
	Username string `json:"username"`
}
