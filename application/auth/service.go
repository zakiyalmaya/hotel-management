package auth

import "github.com/zakiyalmaya/hotel-management/model"

type AuthService interface {
	Login(request *model.AuthRequest) (*model.AuthResponse, error)
	Logout(username string) error
	RefreshAuthToken(username string) (*model.AuthResponse, error)
}
