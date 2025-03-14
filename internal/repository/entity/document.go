package entity

import (
	db "github.com/alanwade2001/go-sepa-db"
)

type Document struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"size:1000000"`
}

func main() {
	persist := db.NewPersist()
	persist.DB.AutoMigrate(&Document{})

}
