package db

import (
	"apm-trace-worker-example/env"
	"database/sql"
	"fmt"

	sqlserver "github.com/denisenkom/go-mssqldb"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

const (
	server = "localhost"
	port   = "1433"
	table  = "product"
)

var Conn *sql.DB

func ConnectDB() error {
	config := fmt.Sprintf("server=%s;port=%s;database=%s;trusted_connection=yes", server, port, table)
	fmt.Println(config)

	sqltrace.Register(env.DB_KIND, &sqlserver.Driver{}, sqltrace.WithServiceName("apm-trace-example-with-gorilla-mux-db"), sqltrace.WithAnalytics(true))

	db, err := sqltrace.Open(env.DB_KIND, config)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Connected to %s-%s", server, table))

	Conn = db

	return nil
}
