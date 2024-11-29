package cmd

import (
	"apm-trace-worker-example/db"
	"apm-trace-worker-example/env"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Worker struct {
	productUseCase string
	repository     string
	dbConn         *sql.DB
}

func NewWorker() *Worker {
	return &Worker{}
}

func (wk *Worker) BootStrap() error {
	env.LoadEnvs()
	err := wk.DbConnInit()
	if err != nil {
		logrus.Error("Erro ao iniciar conexão com banco de dados")
		return err
	}

	logrus.Info("Conexão iniciada com sucesso!")
	return nil
}

func (wk *Worker) DbConnInit() error {
	err := db.ConnectDB()

	if err != nil {
		return err
	}

	return nil
}

func (wk *Worker) Exec() {}
