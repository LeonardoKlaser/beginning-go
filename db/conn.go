package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

// criar conexao
func ConnectDB() (*sql.DB, error) {
	//string connectio
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//abrindo conexao
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//ping pra ver se a conexao abrir com sucesso
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to" + dbname)

	return db, nil

}