package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasamarilla/go-crud/controllers"
	initializers "github.com/lucasamarilla/go-crud/initializer"
	"github.com/lucasamarilla/go-crud/middleware"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConectToDB()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()

	// Cria rotas para os membros
	r.GET("/membro", controllers.MembrosIndex)
	r.GET("/membro/:id", controllers.ShowMembro)
	r.POST("/membro", controllers.MembroCreate)
	r.PUT("/membro/:id", controllers.UpdateMembro)
	r.DELETE("membro/:id", controllers.DeleteMembro)

	//Cuida da valaidação de admin
	r.POST("/signup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
