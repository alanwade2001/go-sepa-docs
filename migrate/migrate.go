package main

import (
	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-docs/internal/repository/entity"
)

func main() {
	persist := db.NewPersist()
	persist.DB.AutoMigrate(&entity.Document{})

}
