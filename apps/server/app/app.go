package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	host   string
	port   string
}

func New(host string, port string) *App {
	router := gin.Default()
	newApp := App{router, host, port}
	return &newApp
}

func (app *App) Start() {
	addr := app.host + ":" + app.port
	fmt.Println("Run")
	app.router.Run(addr)
}
