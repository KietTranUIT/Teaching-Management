package User

type User struct {
	username, password string
}

func (u User) SetUser(usr, pass string) {
	u.username = usr
	u.password = pass
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetPassword() string {
	return u.password
}
