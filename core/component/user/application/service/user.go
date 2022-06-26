package service

import (
	"context"
	"explicit/core/component/user/application/dto"
	"explicit/core/component/user/application/repository"
	"explicit/core/component/user/application/validation"
	"explicit/core/component/user/application/vo"
	"explicit/core/component/user/domain"
)

type UserService struct {
	validation validation.UserValidation
	repository repository.UserRepositoryInterface
}

func (srv *UserService) CreateUser(ctx context.Context, user dto.CreateUser) (*vo.CreateUser, error) {
	if err := srv.validateUserData(user); err != nil {
		return nil, err
	}

	domainUser := domain.User{
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}

	if err := srv.repository.Add(&domainUser); err != nil {
		return nil, err
	}

	rv := vo.CreateUser{
		ID:    domainUser.ID,
		Name:  domainUser.Name,
		Email: domainUser.Email,
		Age:   domainUser.Age,
	}
	return &rv, nil
}

func (srv *UserService) validateUserData(user dto.CreateUser) error {
	if err := srv.validation.ValidateEmail(user.Email); err != nil {
		return err
	}

	return nil
}
