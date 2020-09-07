package database

import (
	"github.com/go-pg/pg/v10"
	"graphqltest/models"
)

type DocumentRepo struct {
	DB *pg.DB
}

func (d *DocumentRepo) GetDocuments() ([]*models.Document, error) {
	var documents []*models.Document
	err := d.DB.Model(&documents).Select()
	return documents, err
}

func (d *DocumentRepo) GetUsersByDocumentUser(userId int) (*models.User, error) {
	user := new(models.User)
	err := d.DB.Model(user).Where("id = ?", userId).Select()
	return user, err
}
