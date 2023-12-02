package user

type UserCreator interface {
	Execute(userDto UserInCreate) error
}
