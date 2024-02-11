package initializers

import (
	"github.com/lucasamarilla/go-crud/models"
)

func SyncDb() {
	DB.AutoMigrate(&models.Membro{})
	DB.AutoMigrate(&models.User{})
}
