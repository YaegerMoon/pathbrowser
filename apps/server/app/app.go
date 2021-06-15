package app

import (
	"fmt"

	"github.com/YaegerMoon/pathbrowser/server/controllers/dir"
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
	dir.New(app.router, "dirs")
	app.router.Run(addr)
}
