package database

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

type Logging interface {
	Printf(ctx context.Context, format string, v ...interface{})
}

const (
	PGADDRESS = ":5432"
	USER      = "postgres"
	PASSWORD  = "postgres"
	DATABASE  = "graph"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	log.Println(string(query))
	return nil
}

func Connect() (connection *pg.DB, ctx context.Context) {
	connection = pg.Connect(&pg.Options{
		Addr:      PGADDRESS,
		User:      USER,
		Password:  PASSWORD,
		Database:  DATABASE,
		TLSConfig: nil,
	})

	log.SetFlags(log.LstdFlags)
	connection.AddQueryHook(dbLogger{})

	ctx = context.Background()
	if err := connection.Ping(ctx); err != nil {
		log.Fatal("DB ERROR", err)
	}

	return connection, ctx
}

func Connect1() (connection *pg.DB) {
	connection = pg.Connect(&pg.Options{
		Addr:      PGADDRESS,
		User:      USER,
		Password:  PASSWORD,
		Database:  DATABASE,
		TLSConfig: nil,
	})

	return connection
}
