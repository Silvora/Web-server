package middleware

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func PasswordToMd5(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(pass), err

}

func ValidatePassword(password string, userPassword string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}
