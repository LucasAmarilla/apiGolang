package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasamarilla/go-crud/controllers"
	initializers "github.com/lucasamarilla/go-crud/initializer"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.MembroCreate)
	r.GET("/", controllers.MembrosIndex)
	r.GET("/:id", controllers.ShowMembro)
	r.PUT("/:id", controllers.UpdateMembro)
	r.DELETE("/:id", controllers.DeleteMembro)

	r.Run()
}
