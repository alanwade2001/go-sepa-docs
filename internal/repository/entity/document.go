package entity

import (
	db "github.com/alanwade2001/go-sepa-db"
	"gorm.io/gorm"
)

type Document struct {
	Model   gorm.Model `gorm:"embedded"`
	Content string     `gorm:"size:1000000"`
}

func main() {
	persist := db.NewPersist()
	persist.DB.AutoMigrate(&Document{})

}
