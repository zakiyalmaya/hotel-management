package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (u *userSvcImpl) Login(request *model.AuthRequest) (*model.AuthResponse, error) {
	response, err := u.repos.UserRepo.GetByUsername(request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("error getting hotelier by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(request.Password)); err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	duration := 15 * time.Minute
	expirationTime := time.Now().Add(duration)
	claims := &model.AuthClaims{
		UserID:   response.ID,
		Username: response.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("hotel-management-secret-key"))
	if err != nil {
		return nil, fmt.Errorf("failed to create token")
	}

	err = u.repos.RedCl.Set(context.Background(), "jwt-token-"+response.Username, tokenString, duration).Err()
	if err != nil {
		log.Println("Failed to store token in Redis:", err.Error())
		return nil, fmt.Errorf("failed to store token in Redis")
	}

	return &model.AuthResponse{
		Username: response.Username,
		Name:     response.Name,
		Token:    tokenString,
	}, nil
}

func (u *userSvcImpl) Logout(username string) error {
	if err := u.repos.RedCl.Del(context.Background(), "jwt-token-"+username).Err(); err != nil {
		log.Println("Failed to delete token from Redis:", err.Error())
		return fmt.Errorf("failed to delete token from Redis")
	}

	return nil
}

func (u *userSvcImpl) RefreshAuthToken(username string) (*model.AuthResponse, error) {
	oldToken, err := u.repos.RedCl.Get(context.Background(), "jwt-token-"+username).Result()
	if err != nil {
		log.Println("Failed to get token from Redis:", err.Error())
		return nil, fmt.Errorf("failed to get token from Redis")
	}

	token, err := jwt.ParseWithClaims(oldToken, &model.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("hotel-management-secret-key"), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Verify token is valid
	if claims, ok := token.Claims.(*model.AuthClaims); ok && token.Valid {
		// Extend expiration time by 15 minutes
		duration := 15 * time.Minute
		claims.ExpiresAt = time.Now().Add(duration).Unix()

		// Generate a new token with updated expiration
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := newToken.SignedString([]byte("hotel-management-secret-key"))
		if err != nil {
			return nil, fmt.Errorf("failed to sign new token: %w", err)
		}

		// Set new token in Redis
		if err := u.repos.RedCl.Set(context.Background(), "jwt-token-"+username, tokenString, duration).Err(); err != nil {
			log.Println("Failed to store token in Redis:", err.Error())
			return nil, fmt.Errorf("failed to store token in Redis")
		}

		return &model.AuthResponse{
			Username: claims.Username,
			Token:    tokenString,
		}, nil
	}

	return nil, fmt.Errorf("token is invalid or claims not valid")
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