package handler

import (
	"net/http"

	"github.com/alanwade2001/go-sepa-docs/internal/service"
	"github.com/alanwade2001/go-sepa-infra/routing"
	"github.com/gin-gonic/gin"
)

type Document struct {
	service *service.Document
}

func NewDocument(service *service.Document, r *routing.Router) *Document {
	document := &Document{
		service: service,
	}

	r.Router.POST("/documents", document.PostDocument)
	r.Router.GET("/documents", document.GetDocument)
	r.Router.GET("/documents/:id", document.GetDocumentByID)

	return document
}

// postInitiation adds an initiations from JSON received in the request body.
func (d *Document) PostDocument(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		return
	}

	newDocument, err := d.service.StoreDocument(string(data))

	code := http.StatusCreated

	if err != nil {
		code = http.StatusInternalServerError
	}

	// Add the new initiations to the slice.
	c.IndentedJSON(code, newDocument)

}

// getDocument responds with the list of all initiationss as JSON.
func (i *Document) GetDocument(c *gin.Context) {
	documents, err := i.service.FindAll()

	code := http.StatusOK

	if err != nil {
		code = http.StatusInternalServerError
	}

	c.IndentedJSON(code, documents)
}

// getDocumentByID locates the initiations whose ID value matches the id
// parameter sent by the client, then returns that initiations as a response.
func (i *Document) GetDocumentByID(c *gin.Context) {
	id := c.Param("id")
	doc, err := i.service.FindByID(id)

	if err != nil {
		message := "Document not found"
		payload := map[string]string{"error": message}
		c.IndentedJSON(http.StatusNotFound, payload)
	} else {
		c.IndentedJSON(http.StatusOK, doc)
	}

}
