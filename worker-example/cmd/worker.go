package cmd

import (
	"apm-trace-worker-example/db"
	"apm-trace-worker-example/dependencies"
	"apm-trace-worker-example/env"
	"apm-trace-worker-example/interfaces"
	"context"
	"fmt"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Worker struct {
	productUseCase interfaces.ProductUseCase
}

func NewWorker() *Worker {
	return &Worker{}
}

func (wk *Worker) BootStrap() error {
	env.LoadEnvs()
	err := wk.DbConnInit()
	if err != nil {
		fmt.Println("Erro ao iniciar conexão com banco de dados")
		return err
	}
	wk.InitiInternalModules()

	fmt.Println("Conexão iniciada com sucesso!")
	return nil
}

func (wk *Worker) DbConnInit() error {
	err := db.ConnectDB()

	if err != nil {
		return err
	}

	return nil
}

func (wk *Worker) InitiInternalModules() {
	wk.productUseCase = dependencies.GetProductUseCase()
}

func (wk *Worker) Exec() {
	span, ctx := tracer.StartSpanFromContext(context.Background(), "worker.exec")
	span.SetTag("tag", "teste")
	defer span.Finish()

	products, err := wk.productUseCase.GetProducts(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(products)
}
