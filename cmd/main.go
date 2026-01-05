package main

import (
	"todolist/config"
	"todolist/internal/handler"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engine      *gin.Engine
	Config      *config.Config
	TodoHandler *handler.TodoHandler
}

func newApp(cfg *config.Config, todoHandler *handler.TodoHandler) *App {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", todoHandler.GetAll)
		v1.GET("/todos/:id", todoHandler.GetByID)
		v1.POST("/todos", todoHandler.Create)
	}

	return &App{
		Engine:      r,
		Config:      cfg,
		TodoHandler: todoHandler,
	}
}

func main() {
	// Llamamos a la función mágica que generará Wire
	app, err := initializeApp()
	if err != nil {
		panic("No se pudo inicializar la aplicación: " + err.Error())
	}

	app.Engine.Run(":" + app.Config.ServerPort)
}
