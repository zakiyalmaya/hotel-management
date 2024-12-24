package user

import "github.com/zakiyalmaya/hotel-management/model"

//go:generate mockery --name=UserService --output=../mocks --outpkg=mocks
type UserService interface {
	Create(request *model.CreateUserRequest) error
	ChangePassword(request *model.ChangePasswordRequest) error
}
