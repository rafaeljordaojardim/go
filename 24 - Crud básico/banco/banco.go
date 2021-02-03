package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver conexão com mysql
)

// Conectar conecta no banco
func Conectar() (*sql.DB, error) {
	url := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", url)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
