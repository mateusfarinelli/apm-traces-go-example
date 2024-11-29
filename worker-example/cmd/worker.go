package cmd

import (
	"apm-trace-worker-example/db"
	"apm-trace-worker-example/dependencies"
	"apm-trace-worker-example/env"
	"apm-trace-worker-example/interfaces"
	"context"
	"fmt"
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
	ctx := context.Background()
	products, err := wk.productUseCase.GetProducts(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(products)
}
