package user

import "github.com/zakiyalmaya/hotel-management/model"

type UserService interface {
	Create(request *model.CreateUserRequest) error
	ChangePassword(request *model.ChangePasswordRequest) error
}
