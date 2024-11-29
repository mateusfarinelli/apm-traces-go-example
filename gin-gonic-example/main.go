package main

import (
	"apm-trace-with-gin-gonic-example/app"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start()
	defer tracer.Stop()

	app := app.NewApp()

	app.Bootstrap()

	app.StartHttpServer()
}