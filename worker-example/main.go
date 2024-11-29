package main

import (
	"apm-trace-worker-example/cmd"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithServiceName("apm-trace-worker-example"),
	)
	defer tracer.Stop()

	wk := cmd.NewWorker()

	wk.BootStrap()

	wk.Exec()

}
