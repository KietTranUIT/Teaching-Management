package User

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	username, password, id, role, email string
}

func (u *User) SetUser(user ...string) {
	u.username = user[0]
	hash := sha256.Sum256([]byte(user[1]))
	u.password = hex.EncodeToString(hash[:])
	u.id = user[2]
	u.role = user[3]
	u.email = user[4]
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetPassword() string {
	return u.password
}

func (u User) GetId() string {
	return u.id
}

func (u User) GetRole() string {
	return u.role
}

func (u User) GetEmail() string {
	return u.email
}
