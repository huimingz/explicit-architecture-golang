package repository

import "explicit/core/component/user/domain"

type UserRepositoryInterface interface {
	Add(user *domain.User) error
	GetOneByID(id int64) (domain.User, error)
	RemoveOneByID(id int64) error
}
