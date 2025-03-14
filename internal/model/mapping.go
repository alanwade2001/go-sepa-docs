package model

import "github.com/alanwade2001/go-sepa-docs/internal/repository/entity"

func ToDocument(doc *entity.Document) *Document {
	document := Document{
		ID:      doc.Model.ID,
		Content: doc.Content,
	}

	return &document
}

func ToDocuments(initns []entity.Document) []*Document {
	documents := []*Document{}

	for _, v := range initns {
		document := ToDocument(&v)
		documents = append(documents, document)
	}

	return documents
}
