package auth

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

type authSvcImpl struct {
	repos *repository.Repositories
}

func NewAuthServiceImpl(repos *repository.Repositories) AuthService {
	return &authSvcImpl{repos: repos}
}

func (a *authSvcImpl) Login(request *model.AuthRequest) (*model.AuthResponse, error) {
	response, err := a.repos.UserRepo.GetByUsername(request.Username)
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

	err = a.repos.RedCl.Set(context.Background(), "jwt-token-"+response.Username, tokenString, duration).Err()
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

func (a *authSvcImpl) Logout(username string) error {
	if err := a.repos.RedCl.Del(context.Background(), "jwt-token-"+username).Err(); err != nil {
		log.Println("Failed to delete token from Redis:", err.Error())
		return fmt.Errorf("failed to delete token from Redis")
	}

	return nil
}

func (a *authSvcImpl) RefreshAuthToken(username string) (*model.AuthResponse, error) {
	oldToken, err := a.repos.RedCl.Get(context.Background(), "jwt-token-"+username).Result()
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
		if err := a.repos.RedCl.Set(context.Background(), "jwt-token-"+username, tokenString, duration).Err(); err != nil {
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
