package user

import "github.com/zakiyalmaya/hotel-management/model"

//go:generate mockery --name=UserRepository --output=../mocks --outpkg=mocks
type UserRepository interface {
	Create(user *model.UserEntity) error
	GetByUsername(username string) (*model.UserEntity, error)
	UpdatePassword(user *model.UserEntity) error
}
