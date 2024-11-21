package user

import "github.com/zakiyalmaya/hotel-management/model"

type UserService interface {
	Create(request *model.CreateUserRequest) error
	Login(request *model.AuthRequest) (*model.AuthResponse, error)
	Logout(username string) error
	RefreshAuthToken(username string) (*model.AuthResponse, error)
	ChangePassword(request *model.ChangePasswordRequest) error
}
