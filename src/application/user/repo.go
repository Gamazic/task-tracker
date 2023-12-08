package user

type SaverUserDto struct {
	Username string
}

type SaveUserUsecase interface {
	SaveCheckFreeUsername(user SaverUserDto) (isFreeUsername bool, err error)
}
