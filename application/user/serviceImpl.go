package user

import (
	"database/sql"
	"fmt"

	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/model"
	"golang.org/x/crypto/bcrypt"
)

type userSvcImpl struct {
	repos *repository.Repositories
}

func NewUserServiceImpl(repos *repository.Repositories) UserService {
	return &userSvcImpl{repos: repos}
}

func (u *userSvcImpl) Create(request *model.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password")
	}

	err = u.repos.UserRepo.Create(&model.UserEntity{
		Name:     request.Name,
		Username: request.Username,
		Password: string(hashedPassword),
		Email:    request.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *userSvcImpl) ChangePassword(request *model.ChangePasswordRequest) error {
	user, err := u.repos.UserRepo.GetByUsername(request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}

		return err
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword)); err != nil {
		return fmt.Errorf("wrong password")
	}

	// compaare old and new password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.NewPassword)); err == nil {
		return fmt.Errorf("new password must be different from old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password")
	}

	user.Password = string(hashedPassword)
	if err := u.repos.UserRepo.UpdatePassword(user); err != nil {
		return err
	}

	return nil
}
