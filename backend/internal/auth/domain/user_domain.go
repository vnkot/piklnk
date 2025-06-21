package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint
	Email    string
	Name     string
	Password string `json:"-"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	} else {
		user.Password = string(hashedPassword)
	}

	return nil
}

func (user *User) IsValidPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
