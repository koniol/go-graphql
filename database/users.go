package database

import (
	"github.com/go-pg/pg/v10"
	"graphqltest/graph/model"
	"graphqltest/models"
	"strconv"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUsers() ([]*models.User, error) {
	var user []*models.User
	err := u.DB.Model(&user).Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) GetUserByID(userId string) (*models.User, error) {
	var user models.User

	value, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	err = u.DB.Model(&user).Where("id = ?", value).Select()

	return &user, err
}

func (u *UserRepo) CreateUser(input model.NewUser) (*models.User, error) {
	user := models.User{
		Email: input.Email,
	}
	_, err := u.DB.Model(&user).Returning("*").Insert()

	return &user, err
}

func (u *UserRepo) GetDocumentsByUserId(userId string) ([]*models.Document, error) {
	var documents []*models.Document
	err := u.DB.Model(&documents).Where("user_id = ?", userId).Select()
	return documents, err
}
