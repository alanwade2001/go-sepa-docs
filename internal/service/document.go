package service

import (
	"github.com/alanwade2001/go-sepa-docs/internal/model"
	"github.com/alanwade2001/go-sepa-docs/internal/repository"
	"github.com/alanwade2001/go-sepa-docs/internal/repository/entity"
)

type Document struct {
	repository *repository.Document
}

func NewDocument(repository *repository.Document) *Document {
	document := &Document{
		repository: repository,
	}

	return document
}

func (d *Document) FindAll() ([]*model.Document, error) {
	docs, err := d.repository.FindAll()

	if err != nil {
		return nil, err
	}

	documents := model.ToDocuments(docs)

	return documents, err
}

func (d *Document) FindByID(id string) (*model.Document, error) {
	doc, err := d.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	document := model.ToDocument(doc)

	return document, err
}

func (d *Document) StoreDocument(data string) (*model.Document, error) {
	newDocument := &model.Document{
		Content: data,
	}

	newDoc := &entity.Document{
		Content: newDocument.Content,
	}

	persisted, err := d.repository.Perist(newDoc)

	if err != nil {
		return nil, err
	}

	newDocument.ID = persisted.Model.ID

	return newDocument, err
}
