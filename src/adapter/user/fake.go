package user

import "tracker_backend/src/domain"

type UserSaverStub struct{}

func (UserSaverStub) SaveIfNotExist(user domain.User) error {
	return nil
}
