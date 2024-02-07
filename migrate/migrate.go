package main

import (
	initializers "github.com/lucasamarilla/go-crud/initializer"
	"github.com/lucasamarilla/go-crud/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Membro{})
}
