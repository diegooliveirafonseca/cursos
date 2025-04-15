package banco

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //drive de conexão
)

func Conectar() (*sql.DB, error) {
	//a string de conexão do mysql é diferente do postgresql
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal("Erro ao conectar com o banco.")
		return nil, erro

	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
