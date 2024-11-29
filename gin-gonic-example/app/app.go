package app

import (
	"apm-trace-with-gin-gonic-example/controller"
	"apm-trace-with-gin-gonic-example/db"
	"apm-trace-with-gin-gonic-example/dependencies"
	"apm-trace-with-gin-gonic-example/env"
	"apm-trace-with-gin-gonic-example/middlewares"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

type App struct {
	ProductControlle controller.ProductController
}

func NewApp() *App {
	return &App{}
}

func(app *App) Bootstrap() {
	env.LoadEnvs()
	app.DbConnInit()
}

func(app *App) DbConnInit() {
	err := db.ConnectDB()

	if err != nil {
		os.Exit(-1)
	}
}

func(app *App) initRouter() *gin.Engine{
	server := gin.Default()
	server.Use(gintrace.Middleware("apm-trace-with-gin-gonic-example"))

	server.GET("/products", dependencies.GetProductController().GetProducts)

	return server
}

func (app *App) StartHttpServer() {
	router := app.initRouter()
	c := middlewares.CORS()

	addr := fmt.Sprintf(":%d", env.HttpPort)
	server := &http.Server{
		Addr:          addr,
		Handler:      c.Handler(router),
		ReadTimeout:  time.Duration(env.HttpReadTimeout) * time.Second,
		WriteTimeout: time.Duration(env.HttpWriteTimeout)* time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Print("Failed to start HTTP server")
		os.Exit(-1)
	}
}