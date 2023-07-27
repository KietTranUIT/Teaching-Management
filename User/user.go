package User

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	username, password string
}

func (u *User) SetUser(usr, pass string) {
	u.username = usr
	hash := sha256.Sum256([]byte(pass))
	u.password = hex.EncodeToString(hash[:])
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetPassword() string {
	return u.password
}
