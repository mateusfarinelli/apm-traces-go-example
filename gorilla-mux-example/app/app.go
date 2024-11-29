package app

import (
	"apm-trace-with-gorilla-mux-example/controller"
	"apm-trace-with-gorilla-mux-example/db"
	"apm-trace-with-gorilla-mux-example/env"
	"apm-trace-with-gorilla-mux-example/middlewares"
	"apm-trace-with-gorilla-mux-example/routes"
	"fmt"
	"net/http"
	"os"
	"time"
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

func (app *App) StartHttpServer() {
	router := routes.NewRouter()
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