package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"pratica.com/api-postgresql/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	connection, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
