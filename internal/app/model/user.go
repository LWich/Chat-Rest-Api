package model

import (
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	Id                int
	Email             string
	EncryptedPassword string
	Password          string
	ExpiresIn         int
	RefreshToken      string
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enp, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enp
	}

	return nil
}

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
}

// ComparerPassword ...
func (u *User) ComparerPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}

func encryptString(password string) (string, error) {
	enp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(enp), nil
}
