package validation

import (
	"errors"
	"explicit/core/port/validation"
	"strings"
)

type UserValidation struct {
	phoneValidator validation.PhoneNumberValidatorInterface
}

func NewUserValidation(phoneValidator validation.PhoneNumberValidatorInterface) *UserValidation {
	return &UserValidation{phoneValidator: phoneValidator}
}

func (u *UserValidation) validateUsername(username string) error {
	username = strings.Trim(username, " ")
	return nil
}

func (u *UserValidation) ValidatePassword(password string) error {
	password = strings.Trim(password, " ")

	if password == "" {
		return errors.New("the password can not be empty")
	}
	if len(password) < 6 {
		return errors.New("the password must be least 6 characters long")
	}
	return nil
}

func (u *UserValidation) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("the email can not be empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("the email should look like a real email address")
	}
	return nil
}

func (u *UserValidation) ValidateMobile(mobile string) error {
	return u.phoneValidator.Validate(mobile, "")
}

func (u *UserValidation) ValidateFullName(fullName string) error {
	if fullName == "" {
		return errors.New("the full name can not be empty")
	}

	return nil
}
