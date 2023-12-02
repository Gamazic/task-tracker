package user_controller

import (
	"fmt"
	"tracker_backend/src/presentation/rest/microframework"
)

type UserRequestModel struct {
	Username string `json:"username"`
}

func (u UserRequestModel) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("%w: non empty field 'username' is required",
			microframework.ValidationErr)
	}
	return nil
}

type UserResponseModel struct {
	Username string `json:"username"`
}
