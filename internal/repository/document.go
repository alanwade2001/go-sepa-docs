package repository

import (
	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-docs/internal/repository/entity"
)

type Document struct {
	database *db.Persist
}

func NewDocument(database *db.Persist) *Document {
	initiation := &Document{
		database: database,
	}

	return initiation
}

func (s *Document) FindAll() ([]entity.Document, error) {
	var initiations []entity.Document
	tx := s.database.DB.Find(&initiations)

	err := tx.Error

	return initiations, err
}

func (s *Document) FindByID(id string) (*entity.Document, error) {
	initiation := &entity.Document{}
	tx := s.database.DB.First(initiation, id)

	err := tx.Error

	return initiation, err
}

func (s *Document) Perist(initn *entity.Document) (*entity.Document, error) {
	tx := s.database.DB.Save(initn)
	err := tx.Error
	return initn, err
}
