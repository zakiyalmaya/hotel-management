package auth

import "github.com/zakiyalmaya/hotel-management/model"

//go:generate mockery --name=AuthService --output=../mocks --outpkg=mocks
type AuthService interface {
	Login(request *model.AuthRequest) (*model.AuthResponse, error)
	Logout(username string) error
	RefreshAuthToken(username string) (*model.AuthResponse, error)
}
