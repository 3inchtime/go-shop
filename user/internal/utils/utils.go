package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GenPassword(passWord string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passWord), bcrypt.DefaultCost)
}

func ValidPassword(passWord, hashedPassWord string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(passWord)); err != nil {
		return false, errors.New("valid password failed")
	}
	return true, nil
}
