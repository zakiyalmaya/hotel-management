package user

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/hotel-management/model"
)

type userRepoImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepoImpl{
		db: db,
	}
}

func (u *userRepoImpl) Create(user *model.UserEntity) error {
	query := "INSERT INTO users (name, username, password, email) VALUES (:name, :username, :password, :email)"
	_, err := u.db.NamedExec(query, user)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (u *userRepoImpl) GetByUsername(username string) (*model.UserEntity, error) {
	user := &model.UserEntity{}
	query := "SELECT id, name, username, password, email, created_at, updated_at FROM users WHERE username = ?"
	
	err := u.db.Get(user, query, username)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userRepoImpl) UpdatePassword(user *model.UserEntity) error {
	query := "UPDATE users SET password = :password WHERE username = :username"
	_, err := u.db.NamedExec(query, user)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}