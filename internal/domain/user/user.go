package user

type User struct {
	Username Username
}

func (u User) Validate() error {
	return u.Username.Validate()
}
