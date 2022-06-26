package validation

import (
	"errors"
	"explicit/core/component/user/application/repository"
	"strings"
)

type UserValidation struct {
	userRepo repository.UserRepositoryInterface
}

func (u *UserValidation) validateUsername(username string) error {
	return nil
}

func (u *UserValidation) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("email is invalid")
	}
	return nil
}
