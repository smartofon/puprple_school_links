package user

import (
	"errors"
	"links/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Find(phone string) (*User, error) {
	user := User{}
	tx := repo.Database.DB.First(&user, "phone=?", phone)
	if tx.Error != nil {
		return nil, errors.New("user not exists")
	}
	return &user, nil
}

func (repo *UserRepository) FindBySession(phone string) (*User, error) {
	user := User{}
	tx := repo.Database.DB.First(&user, "session_id=?", phone)
	if tx.Error != nil {
		return nil, errors.New("session not exists")
	}
	return &user, nil
}
